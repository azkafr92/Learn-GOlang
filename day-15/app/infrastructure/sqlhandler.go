package infrastructure

import (
	"day-15/app/domain"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

var DB *gorm.DB

func InitSqliteDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "app.db")
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&domain.User{}, &domain.Book{})
}
