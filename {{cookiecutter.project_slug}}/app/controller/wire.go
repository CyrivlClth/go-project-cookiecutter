//go:generate go get github.com/google/wire/cmd/wire
//go:generate wire
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package controller

import (
	"github.com/google/wire"

	"{{cookiecutter.project_module_name}}/app/service"
	"{{cookiecutter.project_module_name}}/global/database"
	"{{cookiecutter.project_module_name}}/global/logger"
)

func InitArticleHandler() ArticleHandler {
	wire.Build(NewArticleHandler, service.NewArticleService, database.GetDB, logger.GetInstance)
	return &articleHandler{}
}
