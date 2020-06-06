package service

import (
	"github.com/jinzhu/gorm"

	"{{cookiecutter.project_module_name}}/model"
)

type ArticleFilter struct {
}

type ArticleService interface {
	List(filter ArticleFilter) ([]model.Article, error)
	One(filter ArticleFilter) (model.Article, error)
	Create(article *model.Article) error
	Update(filter ArticleFilter, article model.Article) (model.Article, error)
	Delete(id uint) error
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{db: db}
}

type articleService struct {
	db *gorm.DB
}

func (a *articleService) List(filter ArticleFilter) ([]model.Article, error) {
	return model.ListArticle(a.db)
}

func (a *articleService) One(filter ArticleFilter) (model.Article, error) {
	panic("implement me")
}

func (a *articleService) Create(article *model.Article) error {
	return model.CreateArticle(a.db, article)
}

func (a *articleService) Update(filter ArticleFilter, article model.Article) (model.Article, error) {
	panic("implement me")
}

func (a *articleService) Delete(id uint) error {
	panic("implement me")
}
