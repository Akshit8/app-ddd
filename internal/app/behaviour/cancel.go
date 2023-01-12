package behaviour

import (
	"context"
	"time"

	"github.com/mehdihadeli/go-mediatr"
)

type CreateCancelContextBehaviout struct {
	timeout time.Duration
}

func NewCancel(timeout time.Duration) *CreateCancelContextBehaviout {
	return &CreateCancelContextBehaviout{
		timeout: timeout,
	}
}

func (c *CreateCancelContextBehaviout) Handle(
	ctx context.Context,
	req interface{},
	next mediatr.RequestHandlerFunc,
) (interface{}, error) {
	// Not possible to modify context here
	// ctx, cancel := context.WithTimeout(ctx, c.timeout)
	// defer cancel()

	return next()
}
