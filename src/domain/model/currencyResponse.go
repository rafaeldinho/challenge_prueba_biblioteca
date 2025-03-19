package model

type CurrencyResponse struct {
	Success bool               `json:"success"`
	Quotes  map[string]float64 `json:"quotes"`
}
