package mocks

import (
	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockBookHandler struct {
	mock.Mock
}

func (m *MockBookHandler) ListBooks() *model.BooksResponse {
	mocked := m.Called()
	return mocked.Get(0).(*model.BooksResponse)
}

func (m *MockBookHandler) GetById(val int) *model.BookResponse {
	mocked := m.Called(val)
	return mocked.Get(0).(*model.BookResponse)
}

func (m *MockBookHandler) GetBoxPrice(query *model.BookQuery) *model.BookBoxResponse {
	mocked := m.Called(query)
	if mocked.Get(0) == nil {
		return nil
	}
	return mocked.Get(0).(*model.BookBoxResponse)
}

func (m *MockBookHandler) Create(book *model.Book) (error, bool) {
	mocked := m.Called(book)
	return mocked.Error(0), mocked.Get(1).(bool)
}

type MockBookRepository struct {
	mock.Mock
}

func (u *MockBookRepository) GetAll() ([]model.Book, error) {
	args := u.Called()
	return args.Get(0).([]model.Book), args.Error(1)
}

func (u *MockBookRepository) GetById(ID int) (*model.Book, error){
	args := u.Called(ID)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (u *MockBookRepository) FetchBooks(query *model.BookQuery) (*model.CurrencyResponse, error) {
	args := u.Called(query)
	return args.Get(0).(*model.CurrencyResponse), args.Error(1)
}

func (u *MockBookRepository) Save(book *model.Book) (error, bool) {
	args := u.Called(book)
	return args.Error(0), args.Bool(1)
}

func GetAll() *model.BooksResponse {
	books := []model.Book{
		{ID: 1, Title: "El Quijote", Author: "Miguel de Cervantes", Publisher: "Editorial Planeta", Country: "España", Price: 20.5, Currency: "EUR"},
		{ID: 2, Title: "Cien años de soledad", Author: "Gabriel García Márquez", Publisher: "Sudamericana", Country: "Colombia", Price: 18.0, Currency: "USD"},
		{ID: 3, Title: "1984", Author: "George Orwell", Publisher: "Secker & Warburg", Country: "Reino Unido", Price: 15.0, Currency: "GBP"},
	}
	return &model.BooksResponse{Books: books}
}

func GetBook() *model.BookResponse {
	return &model.BookResponse{
		Book: &model.Book{
			ID:        1,
			Title:     "El Quijote",
			Author:    "Miguel de Cervantes",
			Publisher: "Editorial Planeta",
			Country:   "España",
			Price:     20.5,
			Currency:  "EUR",
		},
	}
}
