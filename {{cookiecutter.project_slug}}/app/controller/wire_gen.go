// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controller

import (
	"{{cookiecutter.project_module_name}}/app/service"
	"{{cookiecutter.project_module_name}}/global/database"
)

// Injectors from wire.go:

func InitArticleHandler() ArticleHandler {
	db := _wireDBValue
	articleService := service.NewArticleService(db)
	controllerArticleHandler := NewArticleHandler(articleService)
	return controllerArticleHandler
}

var (
	_wireDBValue = database.DB
)
