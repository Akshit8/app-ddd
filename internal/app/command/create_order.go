package command

import (
	"context"
	"fmt"
	"time"

	"github.com/Akshit8/app-ddd/internal/domain"
	"github.com/Akshit8/app-ddd/pkg/aggregate"
)

type OrderCreator interface {
	Create(context.Context, *domain.Order) error
}

type CreateOrderRequest struct {
	ID string `validate:"required,uuid4"`
}

type CreateOrderHandler struct {
	creator OrderCreator
}

func NewCreateOrderHandler(creator OrderCreator) *CreateOrderHandler {
	return &CreateOrderHandler{
		creator: creator,
	}
}

func (h CreateOrderHandler) Handle(ctx context.Context, createOrder *CreateOrderRequest) (interface{}, error) {
	order, err := domain.NewOrder(
		domain.OrderID(createOrder.ID),
		domain.NewCustomerID(),
		domain.NewProductID(),
		time.Now,
		domain.Submitted,
		aggregate.NewVersion(),
	)
	if err != nil {
		return nil, fmt.Errorf("creating order: %w", err)
	}

	err = h.creator.Create(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("creating order: %w", err)
	}

	return nil, nil
}
