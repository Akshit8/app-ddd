package behaviour

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/mehdihadeli/go-mediatr"
)

type LatencyBehaviour struct{}

func (LatencyBehaviour) Handle(ctx context.Context, req interface{}, next mediatr.RequestHandlerFunc) (interface{}, error) {
	start := time.Now()

	res, err := next()

	elapsed := time.Since(start)
	log.Printf("Execution for the command (%s) took %s", reflect.TypeOf(req), elapsed)

	return res, err
}
