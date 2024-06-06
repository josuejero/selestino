// internal/repository/user_repository_test.go

package repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/internal/repository"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var userRepo *repository.UserRepository
var userMock sqlmock.Sqlmock

func setupUserRepo(t *testing.T) func() {
	var db *sql.DB
	var err error

	db, userMock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	userRepo = &repository.UserRepository{DB: db}

	return func() {
		db.Close()
	}
}

func TestCreateUser(t *testing.T) {
	teardown := setupUserRepo(t)
	defer teardown()

	user := models.User{
		Username: "testuser",
		Password: "password123",
	}

	userMock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (username, password) VALUES ($1, $2)")).
		WithArgs(user.Username, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := userRepo.CreateUser(user)
	assert.NoError(t, err)

	if err := userMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthenticateUser(t *testing.T) {
	teardown := setupUserRepo(t)
	defer teardown()

	user := models.User{
		Username: "testuser",
		Password: "password123",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	userMock.ExpectQuery("SELECT password FROM users WHERE username = \\$1").
		WithArgs(user.Username).
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword))

	authenticated, err := userRepo.AuthenticateUser(user.Username, user.Password)
	assert.NoError(t, err)
	assert.True(t, authenticated)
}
