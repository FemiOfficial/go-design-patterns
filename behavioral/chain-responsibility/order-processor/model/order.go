package model


type Order struct {
	ID string `json:"id"`
	CustomerID string `json:"customerId"`
	Total float64 `json:"total"`
	Processed bool `json:"processed"`
	ValidationFailed bool `json:"validationFailed"`
	Items []Item `json:"items"`
}