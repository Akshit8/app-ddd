package behaviour

import (
	"context"
	"time"

	"github.com/eyazici90/go-mediator/mediator"
)

type Cancel struct {
	timeout time.Duration
}

func NewCancel(timeout time.Duration) *Cancel {
	return &Cancel{
		timeout: timeout,
	}
}

func (c *Cancel) Process(ctx context.Context, _ mediator.Message, next mediator.Next) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return next(ctx)
}
