package model

type Team struct {
	ID uint64
	Name string
	Shield []byte
	Players []byte
	HistoricalData []HistoricalData
}