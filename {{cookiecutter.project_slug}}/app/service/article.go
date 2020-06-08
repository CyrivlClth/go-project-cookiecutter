package service

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

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

type articleService struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewArticleService(db *gorm.DB, logger *logrus.Logger) ArticleService {
	return &articleService{db: db, logger: logger}
}

func (a *articleService) List(filter ArticleFilter) ([]model.Article, error) {
	a.logger.Info(filter)
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
