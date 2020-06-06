//go:generate wire
//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package controller

import (
	"github.com/google/wire"

	"{{cookiecutter.project_module_name}}/app/service"
	"{{cookiecutter.project_module_name}}/global/database"
)

func InitArticleHandler() ArticleHandler {
	wire.Build(NewArticleHandler, service.NewArticleService, wire.Value(database.DB))
	return &articleHandler{}
}
