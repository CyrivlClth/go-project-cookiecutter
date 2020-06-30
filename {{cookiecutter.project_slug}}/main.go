package main

import (
	"{{cookiecutter.project_module_name}}/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		panic(err)
	}
}
