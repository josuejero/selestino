package repository

import (
	"database/sql"
	"log"

	"github.com/josuejero/selestino/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return err
	}

	_, err = r.DB.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", user.Username, string(hashedPassword), user.Role)
	if err != nil {
		log.Printf("Error inserting user into database: %v", err)
	}
	return err
}

func (r *UserRepository) AuthenticateUser(username, password string) (models.User, bool, error) {
	var user models.User
	row := r.DB.QueryRow("SELECT username, password, role FROM users WHERE username = $1", username)
	err := row.Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		log.Printf("Error fetching user from database: %v", err)
		return user, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("Error comparing password hash: %v", err)
		return user, false, nil
	}

	return user, true, nil
}
