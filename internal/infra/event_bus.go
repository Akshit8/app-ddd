package infra

import (
	"fmt"
	"reflect"
)

type ConsoleBus struct {
}

func NewConsoleBus() *ConsoleBus {
	return &ConsoleBus{}
}

func (ConsoleBus) Publish(event interface{}) {
	fmt.Println("Publishing event:", reflect.TypeOf(event).Name())
}

func (ConsoleBus) PublishAll(events ...interface{}) {
	for _, event := range events {
		fmt.Println("Publishing event:", reflect.TypeOf(event).Name())
	}
}
