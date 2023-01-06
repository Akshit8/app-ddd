package domain

import (
	"time"

	"github.com/Akshit8/app-ddd/pkg/aggregate"
)

type OrderID string

func (o OrderID) String() string {
	return string(o)
}

type OrderStatus int

const (
	Unknow OrderStatus = iota
	Submitted
	Paid
	Shipped
	Cancelled
)

type Order struct {
	aggregate.Root

	id         OrderID
	customerID CustomerID
	productID  ProductID
	status     OrderStatus
	version    aggregate.Version

	createdAt time.Time
}

func (o *Order) Valid() error {
	if o.id == "" || o.customerID == "" || o.productID == "" {
		return ErrInvalidOrder
	}

	return nil
}

func NewOrder(
	id OrderID,
	customerID CustomerID,
	productID ProductID,
	now aggregate.Now,
	status OrderStatus,
	version aggregate.Version,
) (*Order, error) {
	order := Order{
		id:         id,
		customerID: customerID,
		productID:  productID,
		status:     status,
		version:    version,
		createdAt:  now(),
	}

	if err := order.Valid(); err != nil {
		return nil, err
	}

	order.AddEvent(CreatedEvent{id: order.id.String()})

	return &order, nil
}

func (o *Order) Pay() {
	o.status = Paid
	o.AddEvent(PaidEvent{id: o.id.String()})
}

func (o *Order) Cancel() {
	o.status = Cancelled
	o.AddEvent(CancelledEvent{id: o.id.String()})
}

func (o *Order) Ship() error {
	if o.status != Paid {
		return ErrOrderNotPaid
	}

	o.status = Shipped

	return nil
}

func (o *Order) ID() string {
	return o.id.String()
}

func (o *Order) CustomerID() string {
	return o.customerID.String()
}

func (o *Order) ProductID() string {
	return o.productID.String()
}

func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) Version() string {
	return o.version.String()
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}
