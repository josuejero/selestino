// internal/repository/user_repository.go

package repository

import (
	"database/sql"
	"fmt"

	"github.com/josuejero/selestino/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", user.Username, string(hashedPassword), user.Role)
	return err
}

func (r *UserRepository) AuthenticateUser(username, password string) (models.User, bool, error) {
	var user models.User
	row := r.DB.QueryRow("SELECT username, password, role FROM users WHERE username = $1", username)
	err := row.Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		return user, false, err
	}

	// Print the user details and the password being compared
	fmt.Printf("Username: %s, Hashed Password: %s\n", user.Username, user.Password)
	fmt.Printf("Password to Compare: %s\n", password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Print the error from bcrypt
		fmt.Printf("Password Comparison Error: %v\n", err)
		return user, false, nil
	}

	return user, true, nil
}
