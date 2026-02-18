package customer

import (
	"database/sql"
	_ "errors"
	"fmt"
)

type userRepository interface {
	getUserByUsername(username string) (user, error)
	registerUser(username, password, email string) error
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) *mysqlUserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) getUserByUsername(username string) (user, error) {

	query := "SELECT * FROM users WHERE username = ?"

	var user user

	err := r.db.QueryRow(query, username).
		Scan(&user.UserID, &user.UserName, &user.Password, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	return user, nil
}

func (r *mysqlUserRepository) registerUser(username, password, email string) (err error) {
	//If username already exist and username is Unique this will give an error

	_, err = r.db.Exec(
		"INSERT INTO users (username, password, email_address) VALUES (?,?,?)",
		username, password, email,
	)
	fmt.Println("hej")

	return
}
