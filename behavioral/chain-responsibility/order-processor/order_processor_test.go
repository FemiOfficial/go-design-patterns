package orderprocessor

import (
	"godesignpatterns/behavioral/chain-responsibility/order-processor/handlers"
	"godesignpatterns/behavioral/chain-responsibility/order-processor/utils"
	"strings"
	"testing"
)


func TestOrderProcessor_Success(t *testing.T) {

	mockOrder := utils.GetUnProcessedMockOrder()
	validator := &handlers.ValidationHandler{}
	processor := &handlers.ProcessHandler{}

	orderProcessor := new(Processor)

	order, err := orderProcessor.OrderJSONParser(mockOrder)
	if err != nil {
		t.Fatal("order processor should parse order json correctly")
	}

	validator.SetNext(processor)

	processedOrder, processError := orderProcessor.HandleOrder(order, validator)

	if processError != nil {
		t.Fatal("order should not return an error")
	}

	if !processedOrder.Processed {
		t.Fatal("order should be processed here")
	}
}
func TestOrderProcessor_ProcessedValidationFailure(t *testing.T) {

	mockOrder := utils.GetProcessedMockOrder()
	validator := &handlers.ValidationHandler{}
	processor := &handlers.ProcessHandler{}

	orderProcessor := new(Processor)
	
	order, err := orderProcessor.OrderJSONParser(mockOrder)
	if err != nil {
		t.Fatal("order processor should parse order json correctly")
	}

	validator.SetNext(processor)

	if !order.Processed {
		t.Fatal("order should be processed already")
	}
	
	_, processError := orderProcessor.HandleOrder(order, validator)

	if processError == nil {
		t.Fatal("order should return an error")
	}

	if !strings.Contains(processError.Error(), "order has already being processed") {
		t.Fatal("order should return an 'already processed' error")
	}
}
func TestOrderProcessor_InvalidOrderValidationFailure(t *testing.T) {

	mockOrder := utils.GetInvalidMockOrder()
	validator := &handlers.ValidationHandler{}
	processor := &handlers.ProcessHandler{}
	
	orderProcessor := new(Processor)
	
	order, err := orderProcessor.OrderJSONParser(mockOrder)
	if err != nil {
		t.Fatal("order processor should parse order json correctly")
	}
	
	validator.SetNext(processor)
	
	_, processError := orderProcessor.HandleOrder(order, validator)
	
	if processError == nil {
		t.Fatal("order should return an error")
	}
	
	if !strings.Contains(processError.Error(), "invalid order request") {
		t.Fatal("order should return an 'invalid order request' error")
	}

	if order.Processed {
		t.Fatal("order should not be processed here")
	}
}