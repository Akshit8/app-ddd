package domain

import "github.com/google/uuid"

type ProductID string

func (p ProductID) String() string {
	return string(p)
}

func NewProductID() ProductID {
	return ProductID(uuid.New().String())
}
