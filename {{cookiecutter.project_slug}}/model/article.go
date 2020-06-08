package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateArticle(db *gorm.DB, article *Article) error {
	return db.Create(article).Error
}

func ListArticle(db *gorm.DB, where ...interface{}) ([]Article, error) {
	articles := make([]Article, 0)
	err := db.Find(&articles, where...).Error
	return articles, err
}
