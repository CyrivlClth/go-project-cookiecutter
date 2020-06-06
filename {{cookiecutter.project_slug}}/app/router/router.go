package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) error {
	RegisterArticle(engine)
	return nil
}
