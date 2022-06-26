package infrastructure

import (
	"github.com/jinzhu/gorm"
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
	
}
