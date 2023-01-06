package query

import "time"

type (
	OrderView struct {
		ID         string
		CustomerID string
		ProductID  string
		Status     int
		CreatedAt  time.Time
	}

	GetOrderDTO struct {
		OrderView OrderView
	}

	GetOrdersDTO struct {
		Orders []OrderView
	}
)
