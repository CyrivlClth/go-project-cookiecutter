package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{cookiecutter.project_module_name}}/app/service"
	"{{cookiecutter.project_module_name}}/model"
)

type ArticleHandler interface {
	List(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type articleHandler struct {
	s service.ArticleService
}

func NewArticleHandler(s service.ArticleService) ArticleHandler {
	return &articleHandler{s: s}
}

func (a *articleHandler) List(ctx *gin.Context) {
	data, err := a.s.List(service.ArticleFilter{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"list": data})
}

func (a *articleHandler) Retrieve(ctx *gin.Context) {
	panic("implement me")
}

func (a *articleHandler) Update(ctx *gin.Context) {
	panic("implement me")
}

func (a *articleHandler) Create(ctx *gin.Context) {
	var req model.Article
	if err := ctx.Bind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := a.s.Create(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": req})
}

func (a *articleHandler) Delete(ctx *gin.Context) {
	panic("implement me")
}
