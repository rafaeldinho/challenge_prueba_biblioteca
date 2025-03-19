package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/challenge_prueba_biblioteca/src/shared"
	"github.com/challenge_prueba_biblioteca/src/usecase"
)

type bookHandler struct {
	useCase usecase.BookUseCase
}

func NewBookHandler(e *echo.Echo, useCase usecase.BookUseCase) *bookHandler {
	h := &bookHandler{useCase}
	g := e.Group("/books")

	g.GET("", h.listBooks)
	g.GET("/:id", h.getById)
	g.GET("/:id/boxprice", h.getBoxPrice)
	g.POST("", h.create)

	return h
}

func (h *bookHandler) listBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, h.useCase.ListBooks())
}

func (h *bookHandler) getById(c echo.Context) error {
	ID, err := shared.GetIntFromString(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": shared.BadRequestMsg})
	}

	result := h.useCase.GetById(ID)

	if result == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": shared.NotFoundRequestMsg})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *bookHandler) getBoxPrice(c echo.Context) error {
	currency := c.QueryParam("currency")

	if len(currency) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": shared.BadRequestMsg})
	}

	qty, err := shared.GetIntFromString(c.QueryParam("quantity"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": shared.BadRequestMsg})
	}

	ID, err := shared.GetIntFromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": shared.BadRequestMsg})
	}

	result := h.useCase.GetBoxPrice(&model.BookQuery{BookID: ID, CurrencyFrom: currency, Quantity: qty})

	if result == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": shared.NotFoundRequestMsg})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *bookHandler) create(c echo.Context) error {
	var book model.Book
	validate := validator.New()

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": shared.BadRequestMsg})
	}

	if err := validate.Struct(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	result := h.useCase.Create(book)

	if result == nil {
		return c.JSON(http.StatusCreated, map[string]string{"message": shared.OKRequestMsg})
	}

	if result.IsAlreadyCreated {
		return c.JSON(http.StatusConflict, map[string]string{"error": shared.ExistRequestMsg})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
}
