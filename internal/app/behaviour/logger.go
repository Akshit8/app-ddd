package behaviour

import (
	"context"
	"log"
	"reflect"

	"github.com/mehdihadeli/go-mediatr"
)

type LoggingBehaviour struct{}

func (LoggingBehaviour) Handle(ctx context.Context, req interface{}, next mediatr.RequestHandlerFunc) (interface{}, error) {
	log.Println("Starting to Process the command:", reflect.TypeOf(req))

	res, err := next()
	if err != nil {
		return nil, err
	}

	log.Println("Ending to Process the command:", reflect.TypeOf(req))

	return res, nil
}
