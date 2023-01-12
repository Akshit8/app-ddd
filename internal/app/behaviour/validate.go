package behaviour

import (
	"context"
	"log"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/mehdihadeli/go-mediatr"
)

var validate = validator.New()

type ValidationBehaviour struct{}

func (ValidationBehaviour) Handle(ctx context.Context, req interface{}, next mediatr.RequestHandlerFunc) (interface{}, error) {
	log.Println("Validating the command:", reflect.TypeOf(req))

	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	return next()
}
