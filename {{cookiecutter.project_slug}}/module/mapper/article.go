// package mapper
package mapper

import (
	"github.com/jinzhu/gorm"

	"{{cookiecutter.project_module_name}}/module/model"
)

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

func (a *Article) Create(db *gorm.DB, article *model.Article) error {
	return db.Create(article).Error
}

func (a Article) List(db *gorm.DB, pagination *Pagination) ([]model.Article, error) {
	out := make([]model.Article, 0)
	err := pagination.WithPagination(db).Find(&out).Error
	return out, err
}

func (a Article) Retrieve(db *gorm.DB, id uint) (model.Article, error) {
	article := model.Article{}
	err := db.Where("id=?", id).First(&article).Error
	return article, err
}

func (a Article) Delete(db *gorm.DB, id uint) error {
	return db.Where("id=?", id).Delete(&model.Article{}).Error
}
