package handlers

import (
	"log"
	"godesignpatterns/behavioral/chain-responsibility/order-processor/model"
)

type ProcessHandler struct {
	BaseHandler
}

func (processor * ProcessHandler) Handle(order* model.Order) (*model.Order, error) {
	order.Processed = true

	log.Printf("order is processed %s for customer: %s", order.ID, order.CustomerID)

	return order, nil

}