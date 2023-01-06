package app

import (
	"fmt"
	"time"

	"github.com/Akshit8/app-ddd/internal/app/behaviour"
	"github.com/Akshit8/app-ddd/internal/app/command"
	"github.com/Akshit8/app-ddd/internal/app/event"
	"github.com/eyazici90/go-mediator/mediator"
)

type OrderStore struct {
	command.OrderCreator
	command.OrderUpdater
	command.OrderGetter
}

func NewMediator(
	store OrderStore,
	ep event.Publisher,
	timeout time.Duration,
) (*mediator.Mediator, error) {
	m, err := mediator.New(
		mediator.WithBehaviourFunc(behaviour.Logger),
		mediator.WithBehaviourFunc(behaviour.Measure),
		mediator.WithBehaviourFunc(behaviour.Validate),
		mediator.WithBehaviour(behaviour.NewCancel(timeout)),

		mediator.WithHandler(command.CreateOrder{}, command.NewCreateOrderHandler(store)),
		mediator.WithHandler(command.PayOrder{}, command.NewPayOrderHandler(store, store)),
		mediator.WithHandler(command.ShipOrder{}, command.NewShipOrderHandler(store, store, ep)),
		mediator.WithHandler(command.CanceOrder{}, command.NewCancelOrderHandler(store, store)),
	)
	if err != nil {
		return nil, fmt.Errorf("new mediator: %w", err)
	}

	return m, nil
}
