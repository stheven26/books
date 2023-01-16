package repository

import (
	"test/internal/domain/entity"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

type BookRepository interface {
	CreateBook(data entity.Book) (*entity.Book, error)
	GetAllBook(pagination entity.FilterPagination) (*[]entity.Book, error)
}

func InitBookRepository(db *gorm.DB) BookRepository {
	if db == nil {
		panic("initRepository is nil")
	}
	return &bookRepository{db: db}
}

func (b *bookRepository) CreateBook(data entity.Book) (*entity.Book, error) {
	if err := b.db.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *bookRepository) GetAllBook(pagination entity.FilterPagination) (*[]entity.Book, error) {
	data := []entity.Book{}
	offset := (pagination.Page - 1) * pagination.Limit
	if err := b.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Order).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
