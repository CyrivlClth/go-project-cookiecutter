// package router
package router

import (
	"github.com/gin-gonic/gin"

	"{{cookiecutter.project_module_name}}/internal/controller"
)

func New(article *controller.Article) *gin.Engine {
	engine := gin.Default()
	engine.GET("/articles", article.List)
	engine.GET("/articles/:id", article.Retrieve)
	return engine
}
