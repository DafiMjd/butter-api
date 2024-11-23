package pagination

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Docs       interface{} `json:"docs"`
	Limit      int         `json:"limit" query:"limit"`
	Page       int         `json:"page" query:"page"`
	Sort       string      `json:"-" query:"sort"`
	TotalDocs  int64       `json:"totalDocs"`
	TotalPages int         `json:"totalPages"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalDocs int64
	db.Model(value).Count(&totalDocs)

	pagination.TotalDocs = totalDocs
	totalPages := int(math.Ceil(float64(totalDocs) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(pagination.GetOffset()).
			Limit(pagination.GetLimit()).
			Order(pagination.GetSort())
	}
}
