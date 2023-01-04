package domain

import "github.com/google/uuid"

type CustomerID string

func (c CustomerID) String() string {
	return string(c)
}

func NewCustomerID() CustomerID {
	return CustomerID(uuid.New().String())
}
