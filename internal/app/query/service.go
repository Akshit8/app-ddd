package query

import (
	"context"

	"github.com/Akshit8/app-ddd/internal/domain"
)

type OrderQueryStore interface {
	Get(ctx context.Context, id string) (*domain.Order, error)
	GetAll(context.Context) ([]*domain.Order, error)
}

type OrderQueryService struct {
	store OrderQueryStore
}

func NewOrderQueryService(store OrderQueryStore) *OrderQueryService {
	return &OrderQueryService{
		store: store,
	}
}

func (s *OrderQueryService) GetOrder(ctx context.Context, id string) *GetOrderDTO {
	order, err := s.store.Get(ctx, id)
	if err != nil {
		return nil
	}

	orderView := mapTo(order)

	return &GetOrderDTO{
		OrderView: orderView,
	}
}

func (s *OrderQueryService) GetOrders(ctx context.Context) *GetOrdersDTO {
	orders, err := s.store.GetAll(ctx)
	if err != nil {
		return nil
	}

	orderViews := mapAll(orders)

	return &GetOrdersDTO{
		Orders: orderViews,
	}
}
