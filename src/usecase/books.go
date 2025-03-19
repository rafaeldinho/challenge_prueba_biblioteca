package usecase

import (
	"fmt"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/challenge_prueba_biblioteca/src/interface/repository"
)

type BookUseCase interface {
	ListBooks() *model.BooksResponse
	GetById(int) *model.BookResponse
	GetBoxPrice(*model.BookQuery) *model.BookBoxResponse
	Create(model.Book) *model.BookCreated
}

type bookUseCase struct {
	repository repository.BookRepository
}

func NewBookUseCaseUseCase(repository repository.BookRepository) *bookUseCase {
	return &bookUseCase{repository: repository}
}

func (u *bookUseCase) ListBooks() *model.BooksResponse {
	result, err := u.repository.GetAll()

	if err != nil {
		return nil
	}
	return &model.BooksResponse{Books: result}
}

func (u *bookUseCase) GetById(ID int) *model.BookResponse {
	result, err := u.repository.GetById(ID)
	if err != nil {
		return nil
	}
	return &model.BookResponse{Book: result}
}

func (u *bookUseCase) GetBoxPrice(query *model.BookQuery) *model.BookBoxResponse {
	book, err := u.repository.GetById(query.BookID)
	if err != nil {
		return nil
	}
	query.CurrencyTo = book.Currency

	currencies, err := u.repository.FetchBooks(query)
	if err != nil {
		return nil
	}

	key := fmt.Sprintf("%s%s", query.CurrencyFrom, query.CurrencyTo)
	rate, exists := currencies.Quotes[key]
	if !exists {
		return nil
	}

	totalPrice := (book.Price * float64(query.Quantity)) * rate

	return &model.BookBoxResponse{TotalPrice: totalPrice}
}

func (u *bookUseCase) Create(book model.Book) *model.BookCreated {

	if err, exists := u.repository.Save(&book); err != nil {
		return &model.BookCreated{Error: err, IsAlreadyCreated: exists}
	}

	return nil
}
