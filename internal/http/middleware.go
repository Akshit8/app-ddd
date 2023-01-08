package http

import (
	"net/http"
	"time"

	"github.com/Akshit8/app-ddd/internal/domain"
	"github.com/Akshit8/app-ddd/pkg/aggregate"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) useMiddleware() {
	s.echo.Use(middleware.Logger())

	s.echo.Use(middleware.Recover())

	s.echo.Use(middleware.RequestID())

	s.echo.Use(middleware.TimeoutWithConfig(
		middleware.TimeoutConfig{
			Timeout: time.Duration(s.config.Context.Timeout) * time.Second,
		},
	))

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(s.echo)

	s.echo.HTTPErrorHandler = NewHandler(
		defaultHandler.WithMap(
			http.StatusBadRequest,
			aggregate.ErrNotFound,
			domain.ErrInvalidOrder,
			domain.ErrOrderNotPaid,
		),

		defaultHandler.WithMapFunc(
			func(err error) (int, bool) {
				_, ok := err.(validator.ValidationErrors)
				return http.StatusBadRequest, ok
			},
		),
	).Handle()
}
