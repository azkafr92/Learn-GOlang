package interfaces

import (
	"regexp"
	"testing"
	"time"

	"day-15/app/domain"
	"day-15/app/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	now = time.Now()
	users = []domain.User{
		{
			Email:    "xoled72006@runqx.com",
			Password: "ZsIdvdayP72PG4CI",
			Common: domain.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
				DeletedAt: nil,
			},
		},
		{
			Email:    "voxiy70885@tagbert.com",
			Password: "lUUiZtDhOoac2038",
			Common: domain.Common{
				ID:        2,
				CreatedAt: now,
				UpdatedAt: now,
				DeletedAt: nil,
			},
		},
	}
)

func TestUserRepositoryFindAll(t *testing.T) {
	db, mock := mocks.SqliteDBMock()
	
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(users[0].ID, users[0].Email, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt, users[0].DeletedAt).
		AddRow(users[1].ID, users[1].Email, users[1].Password, users[1].CreatedAt, users[1].UpdatedAt, users[1].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).WillReturnRows(rows)

	ur := &UserRepository{db}
	users, err := ur.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	// t.Log(users)

	assert.Len(t, users, 2)
	assert.NoError(t, err)
}

func TestUserRepositorySave(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectExec(query).
		WithArgs(user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ur := &UserRepository{db}
	_, err := ur.Save(&user)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, user.ID, 1)
	assert.EqualValues(t, user.Email, users[0].Email)
	assert.EqualValues(t, user.Password, users[0].Password)
}

func TestUserRepositoryFindById(t *testing.T) {
	db, mock := mocks.SqliteDBMock()
	
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(users[0].ID, users[0].Email, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt, users[0].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((id = ?)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows)

	ur := &UserRepository{db}
	user, err := ur.FindById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)

	assert.EqualValues(t, user.ID, 1)
	assert.EqualValues(t, user.Email, users[0].Email)
	assert.EqualValues(t, user.Password, users[0].Password)
	assert.Nil(t, user.DeletedAt)
}

func TestUserRepositoryUpdateById(t *testing.T) {
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

	ur := &UserRepository{db}
	user, err := ur.UpdateById(uint(users[0].ID), domain.User{Email: users[1].Email, Password: users[1].Password})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)

	assert.EqualValues(t, user.ID, 1)
	assert.EqualValues(t, user.Email, users[1].Email)
	assert.EqualValues(t, user.Password, users[1].Password)
}

func TestUserRepositoryFindByEmailAndPassword(t *testing.T) {
	db, mock := mocks.SqliteDBMock()
	
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(users[0].ID, users[0].Email, users[0].Password, users[0].CreatedAt, users[0].UpdatedAt, users[0].DeletedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((email = ?) AND (password = ?)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(users[0].Email, users[0].Password).
		WillReturnRows(rows)

	ur := &UserRepository{db}
	user, err := ur.FindByEmailAndPassword(users[0].Email, users[0].Password)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)

	assert.EqualValues(t, user.ID, 1)
	assert.EqualValues(t, user.Email, users[0].Email)
	assert.EqualValues(t, user.Password, users[0].Password)
}
