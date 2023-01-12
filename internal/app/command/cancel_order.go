package command

import (
	"context"

	"github.com/Akshit8/app-ddd/internal/domain"
)

type CancelOrderRequest struct {
	OrderID string `validate:"required,uuid4"`
}

type CancelOrderHandler struct {
	orderHandler
}

func NewCancelOrderHandler(orderGetter OrderGetter, orderUpdater OrderUpdater) *CancelOrderHandler {
	return &CancelOrderHandler{
		orderHandler: newOrderHandler(orderGetter, orderUpdater),
	}
}

func (h *CancelOrderHandler) Handle(ctx context.Context, cancelOrderRequest *CancelOrderRequest) (interface{}, error) {
	err := h.update(ctx, cancelOrderRequest.OrderID, func(order *domain.Order) {
		order.Cancel()
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
