package command

import (
	"context"

	"github.com/Akshit8/app-ddd/internal/app/event"
	"github.com/Akshit8/app-ddd/internal/domain"
	"github.com/eyazici90/go-mediator/mediator"
)

type ShipOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (ShipOrder) Key() int {
	return shipCommandKey
}

type ShipOrderHandler struct {
	orderHandler

	eventPublisher event.Publisher
}

func NewShipOrderHandler(
	orderGetter OrderGetter,
	orderUpdater OrderUpdater,
	eventPublisher event.Publisher,
) ShipOrderHandler {
	return ShipOrderHandler{
		orderHandler:   newOrderHandler(orderGetter, orderUpdater),
		eventPublisher: eventPublisher,
	}
}

func (h ShipOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(ShipOrder)
	if !ok {
		return ErrInvalidCommand
	}

	var order *domain.Order
	err := h.updateErr(ctx, cmd.OrderID, func(o *domain.Order) error {
		order = o
		return o.Ship()
	})
	if err != nil {
		return err
	}

	h.eventPublisher.PublishAll(order.Events()...)

	return nil
}
