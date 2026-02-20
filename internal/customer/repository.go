package customer

import (
	"database/sql"
	"fmt"
)

type userRepository interface {
	getUserByUsername(username string) (user, error)
	registerUser(username, password, emailaddress string) error
	userLogin(loginInput, password string) (user, error)
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
		Scan(&user.UserID, &user.UserName, &user.Password, &user.Email)

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
	query := "INSERT INTO user (username, password, email) VALUES (?,?,?)"

	_, err = r.db.Exec(query, username, password, emailaddress)

	return
}

func (r *UserRepository) userLogin(loginInput, password string) (user, error) {

	var user user
	err := r.db.QueryRow(
		"SELECT user_id, username, password, email FROM users WHERE (username = ? OR email = ?) AND password = ?",
		loginInput, loginInput, password).
		Scan(&user.UserID)
	if err != nil {
		fmt.Println("ERROR userLogin in repo:", err)
	}
	return user, err

}
