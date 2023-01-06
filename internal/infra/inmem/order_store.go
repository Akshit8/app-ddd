package inmem

import (
	"sync"

	"github.com/Akshit8/app-ddd/internal/domain"
)

type OrderRepository struct {
	orders map[string]*domain.Order
	mutex  sync.RWMutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*domain.Order),
		mutex:  sync.RWMutex{},
	}
}

func (r *OrderRepository) Save(order *domain.Order) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.orders[order.ID()] = order

	return nil
}

func (r *OrderRepository) Get(id string) (*domain.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	order, ok := r.orders[id]
	if !ok {
		return nil, nil
	}

	return order, nil
}

func (r *OrderRepository) GetAll() ([]*domain.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	orders := make([]*domain.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.orders, id)

	return nil
}
