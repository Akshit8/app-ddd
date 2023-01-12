package app

import (
	"github.com/Akshit8/app-ddd/internal/app/behaviour"
	"github.com/Akshit8/app-ddd/internal/app/command"
	"github.com/Akshit8/app-ddd/internal/app/event"
	"github.com/mehdihadeli/go-mediatr"
)

type OrderStore interface {
	command.OrderCreator
	command.OrderUpdater
	command.OrderGetter
}

func SetupMediator(
	store OrderStore,
	ep event.Publisher,
) {
	createOrder := command.NewCreateOrderHandler(store)
	payOrder := command.NewPayOrderHandler(store, store)
	shipOrder := command.NewShipOrderHandler(store, store, ep)
	cancelOrder := command.NewCancelOrderHandler(store, store)

	mediatr.RegisterRequestPipelineBehaviors(behaviour.LoggingBehaviour{})
	mediatr.RegisterRequestPipelineBehaviors(behaviour.LatencyBehaviour{})
	mediatr.RegisterRequestPipelineBehaviors(behaviour.ValidationBehaviour{})

	mediatr.RegisterRequestHandler[*command.CreateOrderRequest, interface{}](createOrder)
	mediatr.RegisterRequestHandler[*command.PayOrderRequest, interface{}](payOrder)
	mediatr.RegisterRequestHandler[*command.ShipOrderRequest, interface{}](shipOrder)
	mediatr.RegisterRequestHandler[*command.CancelOrderRequest, interface{}](cancelOrder)
}
