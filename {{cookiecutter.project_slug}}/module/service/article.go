// package service
package service

import (
	"github.com/jinzhu/gorm"

	"{{cookiecutter.project_module_name}}/module/mapper"
	"{{cookiecutter.project_module_name}}/module/model"
)

type Article struct {
	db     *gorm.DB
	mapper *mapper.Article
}

func NewArticle(db *gorm.DB, mapper *mapper.Article) *Article {
	return &Article{db: db, mapper: mapper}
}

func (a *Article) List(req *mapper.Pagination) ([]model.Article, error) {
	return a.mapper.List(a.db, req)
}

func (a Article) Retrieve(id uint) (model.Article, error) {
	return a.mapper.Retrieve(a.db, id)
}

func (a Article) Create(req *model.Article) error {
	return a.mapper.Create(a.db, req)
}

func (a Article) Delete(id uint) error {
	return a.mapper.Delete(a.db, id)
}
