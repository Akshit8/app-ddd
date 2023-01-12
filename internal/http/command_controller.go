package http

import (
	"context"
	"net/http"

	"github.com/Akshit8/app-ddd/internal/app/command"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
)

type CommandController struct{}

func (c *CommandController) create(ctx echo.Context) error {
	return handle(
		ctx,
		http.StatusCreated,
		func(ctx context.Context) error {
			_, err := mediatr.Send[*command.CreateOrderRequest, interface{}](ctx, &command.CreateOrderRequest{ID: uuid.New().String()})
			return err
		},
	)
}

func (c *CommandController) cancel(ctx echo.Context) error {
	id := ctx.Param("id")

	return handle(
		ctx,
		http.StatusOK,
		func(ctx context.Context) error {
			_, err := mediatr.Send[*command.CancelOrderRequest, interface{}](ctx, &command.CancelOrderRequest{OrderID: id})
			return err
		},
	)
}

func (c *CommandController) pay(ctx echo.Context) error {
	id := ctx.Param("id")

	return handle(
		ctx,
		http.StatusOK,
		func(ctx context.Context) error {
			_, err := mediatr.Send[*command.PayOrderRequest, interface{}](ctx, &command.PayOrderRequest{OrderID: id})
			return err
		},
	)
}

func (c *CommandController) ship(ctx echo.Context) error {
	id := ctx.Param("id")

	return handle(
		ctx,
		http.StatusOK,
		func(ctx context.Context) error {
			_, err := mediatr.Send[*command.ShipOrderRequest, interface{}](ctx, &command.ShipOrderRequest{OrderID: id})
			return err
		},
	)
}
