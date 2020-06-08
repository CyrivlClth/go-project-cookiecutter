package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"{{cookiecutter.project_module_name}}/model"
)

var db *gorm.DB

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Debug    bool
}

func InitDatabase(config Config) error {
	var err error
	db, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}
	if config.Debug {
		db.LogMode(config.Debug)
	}
	db.AutoMigrate(&model.Article{})
	return nil
}

func GetDB() *gorm.DB {
	return db
}
