// package internal
//+build wireinject

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"{{cookiecutter.project_module_name}}/global/database"
	"{{cookiecutter.project_module_name}}/internal/controller"
	"{{cookiecutter.project_module_name}}/internal/router"
	"{{cookiecutter.project_module_name}}/module"
)

type server struct {
	engine *gin.Engine
	addrs  []string
}

func setup() (*server, error) {
	db, err := injectDB(database.Config{
		Dsn:       ":memory:",
		Debug:     true,
		Singleton: true,
	})
	if err != nil {
		return nil, err
	}

	return injectWeb(db, []string{":8080"})
}

func injectDB(database.Config) (*gorm.DB, error) {
	wire.Build(database.New)

	return &gorm.DB{}, nil
}

func injectWeb(*gorm.DB, []string) (*server, error) {
	wire.Build(
		controller.NewArticle,
		router.New,
		module.ArticleSet,
		wire.Struct(new(server), "*"),
	)

	return &server{}, nil
}
