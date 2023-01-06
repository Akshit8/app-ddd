package http

import (
	"context"
	"net/http"
	"time"

	"github.com/Akshit8/app-ddd/internal/app"
	"github.com/Akshit8/app-ddd/internal/app/command"
	"github.com/Akshit8/app-ddd/internal/app/event"
	"github.com/eyazici90/go-mediator/mediator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CommandController struct {
	sender *mediator.Mediator
}

func NewCommandController(
	store app.OrderStore,
	ep event.Publisher,
	timeout time.Duration,
) (*CommandController, error) {
	sender, err := app.NewMediator(store, ep, timeout)
	if err != nil {
		return nil, err
	}

	return &CommandController{
		sender: sender,
	}, nil
}

func (c *CommandController) create(ctx echo.Context) error {
	return handle(
		ctx,
		http.StatusCreated,
		func(ctx context.Context) error {
			return c.sender.Send(ctx, command.CreateOrder{ID: uuid.New().String()})
		},
	)
}

func (c *CommandController) cancel(ctx echo.Context) error {
	id := ctx.Param("id")

	return handle(
		ctx,
		http.StatusOK,
		func(ctx context.Context) error {
			return c.sender.Send(ctx, command.CanceOrder{OrderID: id})
		},
	)
}

func (c *CommandController) pay(ctx echo.Context) error {
	id := ctx.Param("id")

	return handle(
		ctx,
		http.StatusOK,
		func(ctx context.Context) error {
			return c.sender.Send(ctx, command.PayOrder{OrderID: id})
		},
	)
}

func (c *CommandController) ship(ctx echo.Context) error {
	id := ctx.Param("id")

	return handle(
		ctx,
		http.StatusOK,
		func(ctx context.Context) error {
			return c.sender.Send(ctx, command.ShipOrder{OrderID: id})
		},
	)
}
