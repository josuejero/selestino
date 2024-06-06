// internal/repository/user_repository.go

package repository

import (
	"database/sql"

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
	err := r.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err == nil, err
}
