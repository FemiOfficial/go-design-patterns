package orderprocessor

import (
	"io"
	"encoding/json"
	"godesignpatterns/behavioral/chain-responsibility/order-processor/model"
	"godesignpatterns/behavioral/chain-responsibility/order-processor/handlers"
)

type HandleOrder interface {
	ProcessOrder (*model.Order, error)
}

type Processor struct {}


func (processor *Processor) OrderJSONParser(data io.Reader) (*model.Order, error) {
	order := new(model.Order)
	err := json.NewDecoder(data).Decode(order)
	return order, err
}

func (processor *Processor) HandleOrder(order *model.Order, handler handlers.IHandler)  (*model.Order, error) {
	return handler.Handle(order)
}