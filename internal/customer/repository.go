package customer

import (
	"database/sql"
	"errors"
	"fmt"
)

type userRepository interface {
	getUserByUsername(username string) (user, error)
	registerUser(username, password, emailaddress string) error
}

type UserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) getUserByUsername(username string) (user, error) {

	query := "SELECT * FROM users WHERE username = ?"

	var user user

	err := r.db.QueryRow(query, username).
		Scan(&user.userID, &user.userName, &user.password, &user.emailAddress)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	return user, nil
}

func (r *UserRepository) registerUser(username, password, emailaddress string) (err error) {
	//If username already exist and username is Unique this will give an error
	query := "INSERT INTO user (username, password, emailaddress) VALUES (?,?,?)"

	_, err = r.db.Exec(query, username, password, emailaddress)

	return
}
