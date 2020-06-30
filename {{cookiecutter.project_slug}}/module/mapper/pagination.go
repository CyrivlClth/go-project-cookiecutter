// package mapper
package mapper

import "github.com/jinzhu/gorm"

const (
	DefaultPage     uint = 1
	DefaultPageSize uint = 10
)

type Pagination struct {
	Page     uint `from:"page" json:"page"`
	PageSize uint `from:"pageSize" json:"pageSize"`
}

func (p *Pagination) GetPage() uint {
	if p == nil {
		return DefaultPage
	}
	if p.Page == 0 {
		p.Page = DefaultPage
	}
	return p.Page
}

func (p *Pagination) GetPageSize() uint {
	if p == nil {
		return DefaultPageSize
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultPageSize
	}
	return p.PageSize
}

func (p *Pagination) WithPagination(db *gorm.DB) *gorm.DB {
	if p.GetPage() != 1 {
		db = db.Offset((p.GetPage() - 1) * p.GetPageSize())
	}
	return db.Limit(p.GetPageSize())
}
