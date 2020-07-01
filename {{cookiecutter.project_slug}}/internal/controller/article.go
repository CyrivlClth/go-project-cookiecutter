// package controller
package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"{{cookiecutter.project_module_name}}/module/mapper"
	"{{cookiecutter.project_module_name}}/module/service"
)

type Article struct {
	svc *service.Article
}

func NewArticle(svc *service.Article) *Article {
	return &Article{svc: svc}
}

func (a *Article) List(ctx *gin.Context) {
	req := mapper.Pagination{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	list, err := a.svc.List(&req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"list": list})
}

func (a *Article) Retrieve(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	data, err := a.svc.Retrieve(uint(id))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
