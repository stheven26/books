package book

import (
	"errors"
	"strconv"
	"test/internal/domain/entity"
	"test/internal/domain/repository"
	"test/pkg/constants"
	"time"

	"gorm.io/gorm"
)

type service struct {
	db             *gorm.DB
	bookRepository repository.BookRepository
}

func NewService(db *gorm.DB, repo repository.BookRepository) Service {
	return &service{db: db, bookRepository: repo}
}

func (s *service) CreateBook(req BookRequest) (res DefaultResponse, err error) {
	data := entity.Book{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	resp, err := s.bookRepository.CreateBook(data)
	if err != nil {
		return
	}
	res = DefaultResponse{
		Message: constants.MESSAGE_SUCCESS,
		Status:  constants.STATUS_SUCCESS,
		Data:    resp,
	}
	return
}

func (s *service) GetAllBook(limit string, page string) (res GetAllBookResponse, err error) {
	defaultLimit, defaultPage, order := 25, 1, "id DESC"
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			err = errors.New("Invalid Format")
		}
		defaultLimit = limitInt
	}
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			err = errors.New("Invalid Format")
		}
		defaultPage = pageInt
	}
	resp, err := s.bookRepository.GetAllBook(entity.FilterPagination{
		Limit: defaultLimit,
		Page:  defaultPage,
		Order: order,
	})
	if err != nil {
		return
	}
	pagination := entity.GetPagination(s.db, defaultLimit, defaultPage)
	res = GetAllBookResponse{
		Message: constants.MESSAGE_SUCCESS,
		Status:  constants.STATUS_SUCCESS,
		Data: GetAllBookData{
			Books:      resp,
			Pagination: pagination,
		},
	}
	return
}
