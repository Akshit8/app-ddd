package infra

import (
	"fmt"
	"reflect"
)

type ConsoleEventBus struct {
}

func NewConsoleBus() *ConsoleEventBus {
	return &ConsoleEventBus{}
}

func (ConsoleEventBus) Publish(event interface{}) {
	fmt.Println("Publishing event:", reflect.TypeOf(event).Name(), event)
}

func (ConsoleEventBus) PublishAll(events ...interface{}) {
	for _, event := range events {
		fmt.Println("Publishing event:", reflect.TypeOf(event).Name(), event)
	}
}
