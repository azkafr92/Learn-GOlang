package interfaces

import (
	"bytes"
	"day-14/app/domain"
	"day-14/app/mocks"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserControllerIndex(t *testing.T) {
	// setup database
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(users[0].ID, users[0].Email, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt, users[0].DeletedAt).
		AddRow(users[1].ID, users[1].Email, users[1].Password, users[1].CreatedAt, users[1].UpdatedAt, users[1].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnRows(rows)

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/users")

	// testing
	userController := NewUserController(db)
	if assert.NoError(t, userController.Index(c)) {
		var (
			err  error
			data map[string][]domain.User
			body string = rec.Body.String()
		)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Len(t, data["data"], 2)
		assert.Equal(t, data["data"][0].Email, users[0].Email)
		assert.Equal(t, data["data"][0].Email, users[0].Email)
	}
}

func TestUserControllerStore(t *testing.T) {
	// setup database
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	user := domain.User{
		Email:    users[0].Email,
		Password: users[0].Password,
		Common: domain.Common{
			CreatedAt: users[0].CreatedAt,
			UpdatedAt: users[0].UpdatedAt,
			DeletedAt: users[0].DeletedAt,
		},
	}

	query := regexp.QuoteMeta(`INSERT INTO "users" ("email","password","created_at","updated_at","deleted_at") VALUES (?,?,?,?,?)`)

	// log.Printf("user: %v\n", user)

	mock.ExpectBegin()
	mock.ExpectExec(query).
		WithArgs(user.Email, user.Password, mocks.AnyTime{}, mocks.AnyTime{}, user.DeletedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/users")

	// testing
	userController := NewUserController(db)
	if assert.NoError(t, userController.Store(c)) {
		var (
			err  error
			data map[string]domain.User
			body string = rec.Body.String()
		)
		// log.Printf("body: %v\n", body)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}
		// log.Printf("data: %v\n", data)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, data["data"].Email, users[0].Email)
		assert.Equal(t, data["data"].Password, users[0].Password)
	}
}

func TestUserControllerShow(t *testing.T) {
	// setup database
	db, mock := mocks.SqliteDBMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(users[0].ID, users[0].Email, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt, users[0].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((id = ?)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(users[0].ID).
		WillReturnRows(rows)

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/users")
	c.SetParamNames("id")
	userID := strconv.Itoa(users[0].ID)
	c.SetParamValues(userID)

	// testing
	userController := NewUserController(db)
	if assert.NoError(t, userController.Show(c)) {
		var (
			err  error
			data map[string]domain.User
			body string = rec.Body.String()
		)
		log.Printf("body: %v\n", body)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}
		// log.Printf("data: %v\n", data)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, data["data"].Email, users[0].Email)
		assert.Equal(t, data["data"].Password, users[0].Password)
	}
}

func TestUserControllerEdit(t *testing.T) {
	// setup database
	db, mock := mocks.SqliteDBMock()
	
	defer db.Close()
	
	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(users[0].ID, users[0].Email, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt, users[0].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((id = ?)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "email" = ?, "password" = ?, "created_at" = ?, "updated_at" = ?, "deleted_at" = ? WHERE "users"."deleted_at" IS NULL AND "users"."id" = ?`)).
		WithArgs(users[1].Email, users[1].Password, users[0].CreatedAt, mocks.AnyTime{}, users[0].DeletedAt, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	update := domain.User{Email: users[1].Email, Password: users[1].Password}
	payload, err := json.Marshal(update)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/users")
	c.SetParamNames("id")
	userID := strconv.Itoa(users[0].ID)
	c.SetParamValues(userID)

	// testing
	userController := NewUserController(db)
	if assert.NoError(t, userController.Edit(c)) {
		var (
			err  error
			data map[string]domain.User
			body string = rec.Body.String()
		)
		log.Printf("body: %v\n", body)

		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			t.Fatal(err)
		}
		// log.Printf("data: %v\n", data)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, data["data"].Email, users[1].Email)
		assert.Equal(t, data["data"].Password, users[1].Password)
	}
}
