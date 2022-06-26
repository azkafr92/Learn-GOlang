package interfaces

import (
	"day-14/app/domain"
	"day-14/app/mocks"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var books = []domain.Book{
	{
		Title: "Sound of Fury",
		Common: domain.Common{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
	},
	{
		Title: "Devil's Heir",
		Common: domain.Common{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
	},
}

func TestBookRepositoryFindAll(t *testing.T) {
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at", "deleted_at"}).
		AddRow(books[0].ID, books[0].Title, books[0].CreatedAt, books[0].UpdatedAt, books[0].DeletedAt).
		AddRow(books[1].ID, books[1].Title, books[1].CreatedAt, books[1].UpdatedAt, books[1].DeletedAt)
	
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL`)).WillReturnRows(rows)

	br := &BookRepository{db}
	books, err := br.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	// t.Log(books)

	assert.Len(t, books, 2)
	assert.NoError(t, err)
}

func TestBookRepositorySave(t *testing.T) {
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	book := domain.Book{
		Title: books[0].Title,
		Common: domain.Common{
			CreatedAt: books[0].CreatedAt,
			UpdatedAt: books[0].UpdatedAt,
			DeletedAt: books[0].DeletedAt,
		},
	}

	query := regexp.QuoteMeta(`INSERT INTO "books" ("title","created_at","updated_at","deleted_at") VALUES (?,?,?,?)`)
	
	mock.ExpectBegin()
	mock.ExpectExec(query).
		WithArgs(book.Title, book.CreatedAt, book.UpdatedAt, book.DeletedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	br := &BookRepository{db}
	_, err := br.Save(&book)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, book.ID, 1)
	assert.EqualValues(t, book.Title, books[0].Title)
}

func TestBookRepositoryFindById(t *testing.T) {
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at", "deleted_at"}).
		AddRow(books[0].ID, books[0].Title, books[0].CreatedAt, books[0].UpdatedAt, books[0].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL AND ((id = ?)) ORDER BY "books"."id" ASC LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows)

	br := &BookRepository{db}
	book, err := br.FindById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(book)

	assert.EqualValues(t, book.ID, 1)
	assert.EqualValues(t, book.Title, books[0].Title)
	assert.Nil(t, book.DeletedAt)
}

func TestBookRepositoryUpdateById(t *testing.T) {
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at", "deleted_at"}).
		AddRow(books[0].ID, books[0].Title, books[0].CreatedAt, books[0].UpdatedAt, books[0].DeletedAt)
		
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL AND ((id = ?)) ORDER BY "books"."id" ASC LIMIT 1`)).
	WithArgs(1).
	WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "books" SET "title" = ?, "created_at" = ?, "updated_at" = ?, "deleted_at" = ? WHERE "books"."deleted_at" IS NULL AND "books"."id" = ?`)).
		WithArgs(books[1].Title, books[0].CreatedAt, mocks.AnyTime{}, books[0].DeletedAt, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	br := &BookRepository{db}
	book, err := br.UpdateById(uint(books[0].ID), domain.Book{Title: books[1].Title})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(book)

	assert.EqualValues(t, book.ID, 1)
	assert.EqualValues(t, book.Title, books[1].Title)
	assert.Nil(t, book.DeletedAt)
}
