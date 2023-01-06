package domain

import "errors"

var (
	ErrInvalidOrder = errors.New("invalid order")
	ErrOrderNotPaid = errors.New("order not paid")
)
