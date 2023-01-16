package entity

import (
	"math"

	"gorm.io/gorm"
)

type FilterPagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"order"`
}

type Pagination struct {
	Page        int  `json:"page"`
	TotalPages  int  `json:"totalPages"`
	TotalItems  int  `json:"totalItems"`
	Limit       int  `json:"limit"`
	HasNext     bool `json:"hasNext"`
	HasPrevious bool `json:"hasPrevious"`
}

func GetPagination(db *gorm.DB, limit, page int) Pagination {
	var cnt int64
	db.Model(&Book{}).Count(&cnt)
	totalPages := math.Ceil(float64(cnt) / float64(limit))
	hasNext := false
	if page < int(totalPages) {
		hasNext = true
	}
	return Pagination{
		Page:        page,
		TotalPages:  int(totalPages),
		TotalItems:  int(cnt),
		Limit:       limit,
		HasNext:     hasNext,
		HasPrevious: page > 1,
	}
}
