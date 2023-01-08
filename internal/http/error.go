package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const DefaultHTTPErrorCode = http.StatusInternalServerError

type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type (
	Map      func(err error) (int, bool)
	Mappings []Map
	Option   func(h *Handler)
)

func (m Mappings) find(err error) (int, bool) {
	for _, v := range m {
		if code, ok := v(err); ok {
			return code, ok
		}
	}

	return 0, false
}

type Handler struct {
	httpErrMapping Mappings
	handle         func(err error, c echo.Context)
}

var defaultHandler = &Handler{}

func NewHandler(opts ...Option) *Handler {
	var errHandler Handler
	for _, opt := range opts {
		opt(&errHandler)
	}

	errHandler.setDefaultProblemDetailsHandler()

	return &errHandler
}

func (h *Handler) WithMap(statusCode int, errs ...error) Option {
	return func(h *Handler) {
		h.httpErrMapping = append(h.httpErrMapping, func(err error) (int, bool) {
			for _, e := range errs {
				if errors.Is(err, e) {
					return statusCode, true
				}
			}

			return 0, false
		})
	}
}

func (h *Handler) WithMapFunc(m Map) Option {
	return func(h *Handler) {
		h.httpErrMapping = append(h.httpErrMapping, m)
	}
}

func (h *Handler) Handle() func(err error, c echo.Context) {
	if h.handle == nil {
		return h.handle
	}

	h.setDefaultProblemDetailsHandler()

	return h.handle
}

func (h *Handler) setDefaultProblemDetailsHandler() {
	problemDetailsHandler := func(err error, c echo.Context) {
		code := DefaultHTTPErrorCode

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		mCode, ok := h.httpErrMapping.find(err)
		if ok {
			handleErr(c, prepareProblemDetails(err, "internal-server-error", mCode, c))
			return
		}

		handleErr(c, prepareProblemDetails(err, "application-error", code, c))
	}

	h.handle = problemDetailsHandler
}

func handleErr(c echo.Context, problem Problem) {
	if c.Response().Committed {
		return
	}

	if c.Request().Method == http.MethodHead {
		err := c.NoContent(problem.Status)
		if err != nil {
			c.Logger().Error(err)
		}
		return
	}

	if err := c.JSON(problem.Status, problem); err != nil {
		c.Logger().Error(err)
	}
}

func prepareProblemDetails(err error, typ string, code int, c echo.Context) Problem {
	return Problem{
		Type:     typ,
		Title:    err.Error(),
		Status:   code,
		Detail:   err.Error(),
		Instance: c.Request().RequestURI,
	}
}
