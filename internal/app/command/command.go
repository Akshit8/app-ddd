package command

import (
	"context"
	"fmt"

	"github.com/Akshit8/app-ddd/internal/domain"
	"github.com/Akshit8/app-ddd/pkg/aggregate"
)

const (
	createCommandKey int = iota + 1
	payCommandKey
	cancelCommandKey
	shipCommandKey
)

type (
	OrderGetter interface {
		Get(context.Context, string) (*domain.Order, error)
	}
	OrderUpdater interface {
		Update(context.Context, *domain.Order) error
	}
)

type orderHandler struct {
	orderGetter  OrderGetter
	orderUpdater OrderUpdater
}

func newOrderHandler(
	orderGetter OrderGetter,
	orderUpdater OrderUpdater,
) orderHandler {
	return orderHandler{
		orderGetter:  orderGetter,
		orderUpdater: orderUpdater,
	}
}

func (h orderHandler) update(ctx context.Context, id string, fn func(*domain.Order)) error {
	return h.updateErr(ctx, id, func(order *domain.Order) error {
		fn(order)

		return nil
	})
}

func (h orderHandler) updateErr(ctx context.Context, id string, fn func(*domain.Order) error) error {
	order, err := h.orderGetter.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("getting order: %w", err)
	}

	if order == nil {
		return fmt.Errorf("id: (%s) %s", id, aggregate.ErrNotFound)
	}

	if err := fn(order); err != nil {
		return err
	}

	if err := h.orderUpdater.Update(ctx, order); err != nil {
		return fmt.Errorf("updating order: %w", err)
	}

	return nil
}
