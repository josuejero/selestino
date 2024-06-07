package repository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/josuejero/selestino/internal/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &UserRepository{DB: db}
	user := models.User{Username: "testuser", Password: "testpass", Role: "user"}

	mock.ExpectExec("INSERT INTO users").WithArgs(user.Username, sqlmock.AnyArg(), user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(user)
	assert.NoError(t, err, "Error should be nil")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthenticateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &UserRepository{DB: db}
	user := models.User{Username: "testuser", Password: "testpass", Role: "user"}

	// Hash the password using bcrypt.GenerateFromPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when hashing the password", err)
	}
	user.Password = string(hashedPassword)

	rows := sqlmock.NewRows([]string{"username", "password", "role"}).
		AddRow(user.Username, user.Password, user.Role)
	mock.ExpectQuery("SELECT username, password, role FROM users WHERE username = \\$1").
		WithArgs("testuser").WillReturnRows(rows)

	returnedUser, authenticated, err := repo.AuthenticateUser("testuser", "testpass")
	assert.NoError(t, err, "Error should be nil")

	// Print the hashed password and the result of the comparison
	fmt.Printf("Stored Password: %s\n", user.Password)
	fmt.Printf("Returned User: %+v\n", returnedUser)
	fmt.Printf("Authenticated: %v\n", authenticated)

	if !authenticated {
		t.Errorf("Expected authentication to be true, got false")
	}
	assert.True(t, authenticated, "User should be authenticated")
	assert.Equal(t, user.Username, returnedUser.Username, "Usernames should match")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
