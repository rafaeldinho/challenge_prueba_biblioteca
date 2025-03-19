package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/challenge_prueba_biblioteca/src/shared"
	"github.com/challenge_prueba_biblioteca/src/test/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewbookHandler(t *testing.T) {
	t.Parallel()
	e := echo.New()
	mockUseCase := new(mocks.MockBookHandler)

	t.Run("When call all book list", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		mockUseCase.On("ListBooks").Return(mocks.GetAll())

		bookHandler := NewBookHandler(e, mockUseCase)

		err := bookHandler.listBooks(echoContext)

		var actualResponse *model.BooksResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &actualResponse)

		assert.Equal(t, mocks.GetAll(), actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertCalled(t, "ListBooks")
	})

	t.Run("When ID is valid and book exists", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		expectedBook := mocks.GetBook()

		mockUseCase.On("GetById", 1).Return(expectedBook)

		bookHandler := NewBookHandler(e, mockUseCase)
		err := bookHandler.getById(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var actualBook model.BookResponse
		json.Unmarshal(rec.Body.Bytes(), &actualBook)
		assert.Equal(t, *expectedBook, actualBook)

		mockUseCase.AssertCalled(t, "GetById", 1)
	})

	t.Run("When ID is invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/abc", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("abc")

		bookHandler := NewBookHandler(e, mockUseCase)
		err := bookHandler.getById(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		expectedResponse := map[string]string{"error": shared.BadRequestMsg}
		var actualResponse map[string]string
		json.Unmarshal(rec.Body.Bytes(), &actualResponse)
		assert.Equal(t, expectedResponse, actualResponse)
	})

	t.Run("When book does not exist", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/99", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("99")

		mockUseCase.On("GetById", 99).Return((*model.BookResponse)(nil))

		bookHandler := NewBookHandler(e, mockUseCase)
		err := bookHandler.getById(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)

		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, shared.NotFoundRequestMsg, response["error"])
		mockUseCase.AssertCalled(t, "GetById", 99)
	})
}

func TestCreateHandler(t *testing.T) {
	t.Run("When book is created successfully", func(t *testing.T) {
		e := echo.New()
		mockUseCase := new(mocks.MockBookHandler)
		bookHandler := NewBookHandler(e, mockUseCase)
		mockBook := mocks.GetBook().Book
		bookJSON, _ := json.Marshal(mockBook)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockUseCase.On("Create", mock.AnythingOfType("*model.Book")).Return(nil, false)

		err := bookHandler.create(ctx)

		assert.NoError(t, err)
	})

	t.Run("When request has invalid data (Bind fails)", func(t *testing.T) {
		e := echo.New()
		mockUseCase := new(mocks.MockBookHandler)
		bookHandler := NewBookHandler(e, mockUseCase)
		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err := bookHandler.create(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "`+shared.BadRequestMsg+`"}`, rec.Body.String())
	})

	t.Run("When validation fails (missing required fields)", func(t *testing.T) {
		e := echo.New()
		mockUseCase := new(mocks.MockBookHandler)
		bookHandler := NewBookHandler(e, mockUseCase)
		mockBook := model.Book{}
		bookJSON, _ := json.Marshal(mockBook)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err := bookHandler.create(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
	})

	t.Run("When book already exists", func(t *testing.T) {
		e := echo.New()
		mockUseCase := new(mocks.MockBookHandler)
		bookHandler := NewBookHandler(e, mockUseCase)
		mockBook := mocks.GetBook().Book
		bookJSON, _ := json.Marshal(mockBook)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockUseCase.On("Create", mock.AnythingOfType("*model.Book")).Return(nil, true)

		err := bookHandler.create(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.JSONEq(t, `{"error": "`+shared.ExistRequestMsg+`"}`, rec.Body.String())
		mockUseCase.AssertCalled(t, "Create", mock.AnythingOfType("*model.Book"))
	})

	t.Run("When an unexpected error occurs", func(t *testing.T) {
		e := echo.New()
		mockUseCase := new(mocks.MockBookHandler)
		bookHandler := NewBookHandler(e, mockUseCase)
		mockBook := mocks.GetBook().Book
		bookJSON, _ := json.Marshal(mockBook)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockUseCase.On("Create", mock.AnythingOfType("*model.Book")).Return(errors.New("unexpected error"), false)

		err := bookHandler.create(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "unexpected error"}`, rec.Body.String())
	})

}

func TestGetBoxPrice(t *testing.T) {
	e := echo.New()
	mockUseCase := new(mocks.MockBookHandler)
	bookHandler := NewBookHandler(e, mockUseCase)

	t.Run("When query is successful", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1/price?currency=USD&quantity=5", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		expectedResponse := &model.BookBoxResponse{
			TotalPrice: 100.0,
		}

		mockUseCase.On("GetBoxPrice", &model.BookQuery{BookID: 1, CurrencyFrom: "USD", Quantity: 5}).Return(expectedResponse)

		err := bookHandler.getBoxPrice(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertCalled(t, "GetBoxPrice", &model.BookQuery{BookID: 1, CurrencyFrom: "USD", Quantity: 5})
	})

	t.Run("When currency is missing", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1/price?quantity=5", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := bookHandler.getBoxPrice(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "`+shared.BadRequestMsg+`"}`, rec.Body.String())
	})

	t.Run("When quantity is invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1/price?currency=USD&quantity=abc", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := bookHandler.getBoxPrice(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "`+shared.BadRequestMsg+`"}`, rec.Body.String())
	})

	t.Run("When book ID is invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/xyz/price?currency=USD&quantity=5", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("xyz")

		err := bookHandler.getBoxPrice(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "`+shared.BadRequestMsg+`"}`, rec.Body.String())
	})
}
