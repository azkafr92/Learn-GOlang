package mocks

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func SqliteDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", dbMock)
	if err != nil {
		panic(err)
	}

	return db, mock
}
