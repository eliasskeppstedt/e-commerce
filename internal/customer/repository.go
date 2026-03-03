package customer

import (
	"database/sql"
	"fmt"
)

type userRepository interface {
	registerUser(username,
		password,
		email,
		first_name,
		last_name,
		address,
		zip_code,
		phone_number string) error
	userLogin(loginInput, password string) (int, bool, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) registerUser(username,
	password,
	email,
	first_name,
	last_name,
	address,
	zip_code,
	phone_number string) (err error) {
	//If username already exist and username is Unique this will give an error
	query := "INSERT INTO users (username, password, email, first_name, last_name, address, zip_code, phone_number) VALUES (?,?,?,?,?,?,?,?)"

	_, err = r.db.Exec(query, username, password, email, first_name, last_name, address, zip_code, phone_number)

	return
}

func (r *UserRepository) userLogin(loginInput, password string) (int, bool, error) {
	var is_admin bool
	var userID int
	err := r.db.QueryRow(
		"SELECT user_id, username, password, email, is_admin FROM users WHERE (username = ? OR email = ?) AND password = ?",
		loginInput, loginInput, password).
		Scan(&userID, &is_admin)
	if err != nil {
		fmt.Println("ERROR userLogin in repo:", err)
	}
	return userID, is_admin, err

}
