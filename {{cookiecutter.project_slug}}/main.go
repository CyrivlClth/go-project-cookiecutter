package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"{{cookiecutter.project_module_name}}/app/router"
	"{{cookiecutter.project_module_name}}/global/config"
)

func main() {
	if err := config.InitConfig("config/config.toml"); err != nil {
		panic(err)
	}

	engine := gin.Default()
	if err := router.InitRouter(engine); err != nil {
		panic(err)
	}
	fmt.Println("http://localhost:8080")
	if err := engine.Run(); err != nil {
		panic(err)
	}
}
