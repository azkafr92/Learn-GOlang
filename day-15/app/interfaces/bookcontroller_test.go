package interfaces

import (
	"bytes"
	"day-15/app/domain"
	"day-15/app/mocks"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBookControllerIndex(t *testing.T) {
	// setup database
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at", "deleted_at"}).
		AddRow(books[0].ID, books[0].Title, books[0].CreatedAt, books[0].UpdatedAt, books[0].DeletedAt).
		AddRow(books[1].ID, books[1].Title, books[1].CreatedAt, books[1].UpdatedAt, books[1].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL`)).WillReturnRows(rows)

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/books")

	// testing
	bookController := NewBookController(db)
	if assert.NoError(t, bookController.Index(c)) {
		var (
			err  error
			data map[string][]domain.Book
			body string = rec.Body.String()
		)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Len(t, data["data"], 2)
		assert.Equal(t, data["data"][0].Title, books[0].Title)
	}
}

func TestBookControllerStore(t *testing.T) {
	// setup database
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
		WithArgs(book.Title, mocks.AnyTime{}, mocks.AnyTime{}, book.DeletedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/books")

	// testing
	bookController := NewBookController(db)
	if assert.NoError(t, bookController.Store(c)) {
		var (
			err  error
			data map[string]domain.Book
			body string = rec.Body.String()
		)
		// log.Printf("body: %v\n", body)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}
		// log.Printf("data: %v\n", data)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, data["data"].Title, books[0].Title)
	}
}

func TestBookControllerShow(t *testing.T) {
	// setup database
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at", "deleted_at"}).
		AddRow(books[0].ID, books[0].Title, books[0].CreatedAt, books[0].UpdatedAt, books[0].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL AND ((id = ?)) ORDER BY "books"."id" ASC LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows)

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/books")
	c.SetParamNames("id")
	bookID := strconv.Itoa(books[0].ID)
	c.SetParamValues(bookID)

	// testing
	bookController := NewBookController(db)
	if assert.NoError(t, bookController.Show(c)) {
		var (
			err  error
			data map[string]domain.Book
			body string = rec.Body.String()
		)

		// log.Printf("body: %v\n", body)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}

		// log.Printf("data: %v\n", data)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, data["data"].Title, books[0].Title)
	}
}

func TestBookControllerEdit(t *testing.T) {
	// setup database
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

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	update := domain.Book{Title: books[1].Title}
	payload, err := json.Marshal(update)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/books")
	c.SetParamNames("id")
	bookID := strconv.Itoa(books[0].ID)
	c.SetParamValues(bookID)

	// testing
	bookController := NewBookController(db)
	if assert.NoError(t, bookController.Edit(c)) {
		var (
			err  error
			data map[string]domain.Book
			body string = rec.Body.String()
		)
		
		// log.Printf("body: %v\n", body)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}
		
		// log.Printf("data: %v\n", data)
		
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, data["data"].Title, books[1].Title)
	}
}
