package model

type Book struct {
	ID        int     `json:"id" validate:"required" bson:"_id"`
	Title     string  `json:"title" validate:"required"`
	Author    string  `json:"author" validate:"required"`
	Publisher string  `json:"publisher" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency" validate:"required,len=3"`
}
