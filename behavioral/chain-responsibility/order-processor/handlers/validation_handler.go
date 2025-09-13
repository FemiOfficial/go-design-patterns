package handlers

import (
	"fmt"
	"godesignpatterns/behavioral/chain-responsibility/order-processor/model"
	"log"
)

type ValidationHandler struct {
	BaseHandler
}

func (validator * ValidationHandler) Handle(order* model.Order) (*model.Order, error) {
	if (order.Processed) { 
		return order, fmt.Errorf("order has already being processed %s", order.ID)
	}

	if (order.ID == "" || order.CustomerID == ""  || len(order.Items) < 1 || order.Total < 1) {
		return order, fmt.Errorf("invalid order request")
	}

	log.Printf("[all validations passed]: order: %s", order.ID)
	return validator.BaseHandler.Handle(order)

}
