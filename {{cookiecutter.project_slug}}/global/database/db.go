package database

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	ins  *gorm.DB
	once sync.Once
)

type Config struct {
	Dsn       string
	Debug     bool
	Singleton bool
}

func New(config Config) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", config.Dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(config.Debug)
	if config.Singleton {
		once.Do(func() {
			ins = db
		})
	}
	return db, nil
}

func GetInstance() *gorm.DB {
	return ins
}
