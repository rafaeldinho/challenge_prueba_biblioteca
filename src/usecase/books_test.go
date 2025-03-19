package usecase

import (
	"errors"
	"testing"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/challenge_prueba_biblioteca/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListBooks(t *testing.T) {
	t.Run("When books exist", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)
		mockBooks := []model.Book{{ID: 1, Title: "Book 1"}, {ID: 2, Title: "Book 2"}}
		mockRepo.On("GetAll").Return(mockBooks, nil)

		result := useCase.ListBooks()

		assert.NotNil(t, result)
		assert.Len(t, result.Books, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("When no books exist", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)
		mockRepo.On("GetAll").Return([]model.Book{}, errors.New("database error"))

		result := useCase.ListBooks()

		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("When book is found", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)

		mockBook := &model.Book{ID: 1, Title: "Book 1"}
		mockRepo.On("GetById", 1).Return(mockBook, nil)

		result := useCase.GetById(1)

		assert.NotNil(t, result)
		assert.Equal(t, mockBook.Title, result.Book.Title)
		mockRepo.AssertExpectations(t)
	})

	t.Run("When book is not found", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)
		mockRepo.On("GetById", 2).Return(&model.Book{}, errors.New("not found"))

		result := useCase.GetById(2)

		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetBoxPrice(t *testing.T) {

	t.Run("When book exists and currency rate is available", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)
		mockBook := &model.Book{ID: 1, Price: 10.0, Currency: "USD"}
		mockRepo.On("GetById", 1).Return(mockBook, nil)

		mockCurrencies := &model.CurrencyResponse{Quotes: map[string]float64{"USDUSD": 1.0}}
		mockRepo.On("FetchBooks", mock.Anything).Return(mockCurrencies, nil)

		query := &model.BookQuery{BookID: 1, CurrencyFrom: "USD", Quantity: 2}
		result := useCase.GetBoxPrice(query)

		assert.NotNil(t, result)
		assert.Equal(t, 20.0, result.TotalPrice)
		mockRepo.AssertExpectations(t)
	})

	t.Run("When book does not exist", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)
		mockRepo.On("GetById", 2).Return(&model.Book{}, errors.New("not found"))

		query := &model.BookQuery{BookID: 2, CurrencyFrom: "USD", Quantity: 1}
		result := useCase.GetBoxPrice(query)

		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("When book is created successfully", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)
		mockBook := &model.Book{Title: "New Book"}
		mockRepo.On("Save", mockBook).Return(nil, true)

		err, created := useCase.Create(mockBook)

		assert.NoError(t, err)
		assert.True(t, created)
		mockRepo.AssertExpectations(t)
	})

	t.Run("When book creation fails", func(t *testing.T) {
		mockRepo := new(mocks.MockBookRepository)
		useCase := NewBookUseCaseUseCase(mockRepo)
		mockBook := &model.Book{Title: "Failed Book"}
		mockRepo.On("Save", mockBook).Return(errors.New("database error"), false)

		err, created := useCase.Create(mockBook)

		assert.Error(t, err)
		assert.False(t, created)
		mockRepo.AssertExpectations(t)
	})
}
