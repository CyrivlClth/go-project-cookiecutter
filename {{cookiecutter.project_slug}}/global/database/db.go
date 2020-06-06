package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}

func Migrate(models ...interface{}) error {
	return DB.AutoMigrate(models...).Error
}
