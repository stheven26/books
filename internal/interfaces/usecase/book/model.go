package book

import (
	"test/internal/domain/entity"
	"time"
)

type (
	DefaultResponse struct {
		Message string      `json:"message"`
		Status  string      `json:"status"`
		Data    interface{} `json:"data"`
	}

	BookRequest struct {
		ID          int64      `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		CreatedAt   time.Time  `json:"createdAt"`
		UpdatedAt   time.Time  `json:"updatedAt"`
		DeletedAt   *time.Time `json:"deletedAt"`
	}

	GetAllBookData struct {
		Books      *[]entity.Book
		Pagination entity.Pagination
	}

	GetAllBookResponse struct {
		Message string         `json:"message"`
		Status  string         `json:"status"`
		Data    GetAllBookData `json:"data"`
	}
)
