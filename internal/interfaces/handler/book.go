package handler

import (
	"net/http"
	"test/internal/interfaces/usecase/book"

	"github.com/labstack/echo/v4"
)

type handler struct {
	bookService book.Service
}

type BookHandler interface {
	CreateBook(c echo.Context) error
	GetAllBook(c echo.Context) error
}

func InitBookHandler(service book.Service) BookHandler {
	return &handler{bookService: service}
}

func (h *handler) CreateBook(c echo.Context) (err error) {
	var req book.BookRequest
	if err = c.Bind(&req); err != nil {
		return
	}
	res, err := h.bookService.CreateBook(req)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *handler) GetAllBook(c echo.Context) (err error) {
	limit, page := c.QueryParam("limit"), c.QueryParam("page")
	res, err := h.bookService.GetAllBook(limit, page)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)

}
