package handlers

import (
	"godesignpatterns/behavioral/chain-responsibility/order-processor/model"
)


type IHandler interface {
	SetNext(handler IHandler) IHandler
	Handle(order *model.Order) (*model.Order, error)
}

type BaseHandler struct {
	next IHandler
}

func (baseHandler * BaseHandler) SetNext(next IHandler) IHandler {
	baseHandler.next = next
	return next
}

func (baseHandler *BaseHandler) Handle(order *model.Order) (*model.Order, error) {
	if (baseHandler.next != nil) {
		result, err := baseHandler.next.Handle(order)
		return result, err
	}

	return order, nil
}