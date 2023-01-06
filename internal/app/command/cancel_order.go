package command

import (
	"context"

	"github.com/Akshit8/app-ddd/internal/domain"
	"github.com/eyazici90/go-mediator/mediator"
)

type CanceOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (CanceOrder) Key() int {
	return cancelCommandKey
}

type CancelOrderHandler struct {
	orderHandler
}

func NewCancelOrderHandler(orderGetter OrderGetter, orderUpdater OrderUpdater) CancelOrderHandler {
	return CancelOrderHandler{
		orderHandler: newOrderHandler(orderGetter, orderUpdater),
	}
}

func (h CancelOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CanceOrder)
	if !ok {
		return ErrInvalidCommand
	}

	return h.update(ctx, cmd.OrderID, func(order *domain.Order) {
		order.Cancel()
	})
}
