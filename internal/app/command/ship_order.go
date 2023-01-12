package command

import (
	"context"

	"github.com/Akshit8/app-ddd/internal/app/event"
	"github.com/Akshit8/app-ddd/internal/domain"
)

type ShipOrderRequest struct {
	OrderID string `validate:"required,uuid4"`
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

func (h ShipOrderHandler) Handle(ctx context.Context, shipOrderRequest *ShipOrderRequest) (interface{}, error) {
	var order *domain.Order
	err := h.updateErr(ctx, shipOrderRequest.OrderID, func(o *domain.Order) error {
		order = o
		return o.Ship()
	})
	if err != nil {
		return nil, err
	}

	h.eventPublisher.PublishAll(order.Events()...)

	return nil, nil
}
