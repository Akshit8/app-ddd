package command

import (
	"context"

	"github.com/Akshit8/app-ddd/internal/domain"
	"github.com/eyazici90/go-mediator/mediator"
)

type PayOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (PayOrder) Key() int {
	return payCommandKey
}

type PayOrderHandler struct {
	orderHandler
}

func NewPayOrderHandler(orderGetter OrderGetter, orderUpdater OrderUpdater) PayOrderHandler {
	return PayOrderHandler{
		orderHandler: newOrderHandler(orderGetter, orderUpdater),
	}
}

func (h PayOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(PayOrder)
	if !ok {
		return ErrInvalidCommand
	}

	return h.update(ctx, cmd.OrderID, func(order *domain.Order) {
		order.Pay()
	})
}
