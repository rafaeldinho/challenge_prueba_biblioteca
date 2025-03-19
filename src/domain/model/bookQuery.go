package model

type BookQuery struct {
	BookID       int    `query:"id" validate:""`
	CurrencyFrom string `query:"currency" validate:"omitempty"`
	CurrencyTo   string `query:"currencyTo"`
	Quantity     int    `query:"quantity" validate:"omitempty,gte=0"`
}
