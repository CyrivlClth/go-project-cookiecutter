// package service
package module

import (
	"github.com/google/wire"

	"{{cookiecutter.project_module_name}}/module/mapper"
	"{{cookiecutter.project_module_name}}/module/service"
)

var ArticleSet = wire.NewSet(service.NewArticle, mapper.NewArticle)
