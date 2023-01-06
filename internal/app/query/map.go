package query

import "github.com/Akshit8/app-ddd/internal/domain"

func mapTo(order *domain.Order) OrderView {
	return OrderView{
		ID:         order.ID(),
		CustomerID: order.CustomerID(),
		ProductID:  order.ProductID(),
		Status:     int(order.Status()),
		CreatedAt:  order.CreatedAt(),
	}
}

func mapAll(orders []*domain.Order) []OrderView {
	views := make([]OrderView, len(orders))
	for i, order := range orders {
		views[i] = mapTo(order)
	}

	return views
}
