package main

import (
	"github.com/gin-gonic/gin"

	"{{cookiecutter.project_module_name}}/app/router"
	_ "{{cookiecutter.project_module_name}}/global/database"
)

func main() {
	engine := gin.Default()
	if err := router.InitRouter(engine); err != nil {
		panic(err)
	}
	if err := engine.Run(); err != nil {
		panic(err)
	}
}
