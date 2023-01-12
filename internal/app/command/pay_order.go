package command

import (
	"context"

	"github.com/Akshit8/app-ddd/internal/domain"
)

type PayOrderRequest struct {
	OrderID string `validate:"required,uuid4"`
}

type PayOrderHandler struct {
	orderHandler
}

func NewPayOrderHandler(orderGetter OrderGetter, orderUpdater OrderUpdater) PayOrderHandler {
	return PayOrderHandler{
		orderHandler: newOrderHandler(orderGetter, orderUpdater),
	}
}

func (h PayOrderHandler) Handle(ctx context.Context, payOrderRequest *PayOrderRequest) (interface{}, error) {
	err := h.update(ctx, payOrderRequest.OrderID, func(order *domain.Order) {
		order.Pay()
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
