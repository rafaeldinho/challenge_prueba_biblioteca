package model

type BookCreated struct {
	IsAlreadyCreated bool
	Error            error
}

type BooksResponse struct {
	Books []Book `json:"books"`
}

type BookResponse struct {
	Book *Book `json:"book"`
}

type BookBoxResponse struct {
	TotalPrice float64 `json:"totalPrice"`
}
