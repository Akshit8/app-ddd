package behaviour

import (
	"context"
	"log"
	"reflect"

	"github.com/eyazici90/go-mediator/mediator"
)

func Logger(ctx context.Context, msg mediator.Message, next mediator.Next) error {
	log.Println("Starting to Process the command:", reflect.TypeOf(msg))

	if err := next(ctx); err != nil {
		return err
	}

	log.Println("Ending to Process the command:", reflect.TypeOf(msg))

	return nil
}
