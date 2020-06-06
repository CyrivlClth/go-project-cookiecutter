package router

import (
	"github.com/gin-gonic/gin"

	"{{cookiecutter.project_module_name}}/app/controller"
)

func RegisterArticle(engine *gin.Engine) {
	r := engine.Group("/resource")
	h := controller.InitArticleHandler()
	r.GET("/article", h.List)
	r.POST("/article", h.Create)
	r.GET("/article/:id", h.Retrieve)
	r.PUT("/article/:id", h.Update)
	r.DELETE("/article/:id", h.Delete)
}
