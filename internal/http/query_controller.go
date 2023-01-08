package http

import (
	"context"
	"net/http"

	"github.com/Akshit8/app-ddd/internal/app/query"
	"github.com/labstack/echo/v4"
)

type QueryService interface {
	GetOrder(ctx context.Context, id string) *query.GetOrderDTO
	GetOrders(ctx context.Context) *query.GetOrdersDTO
}

type QueryController struct {
	service QueryService
}

func NewQueryController(queryService QueryService) QueryController {
	return QueryController{
		service: queryService,
	}
}

func (c *QueryController) getOrder(ctx echo.Context) error {
	id := ctx.Param("id")

	return handleR(
		ctx,
		http.StatusOK,
		func(ctx context.Context) (interface{}, error) {
			return c.service.GetOrder(ctx, id), nil
		},
	)
}

func (c *QueryController) getOrders(ctx echo.Context) error {
	return handleR(
		ctx,
		http.StatusOK,
		func(ctx context.Context) (interface{}, error) {
			return c.service.GetOrders(ctx), nil
		},
	)
}
