// package config
package config

import (
	"github.com/BurntSushi/toml"

	"{{cookiecutter.project_module_name}}/global/database"
)

var Config = struct {
	Database database.Config
}{}

func InitConfig(path string) error {
	if _, err := toml.DecodeFile(path, &Config); err != nil {
		return err
	}

	if err := database.InitDatabase(Config.Database); err != nil {
		return err
	}

	return nil
}
