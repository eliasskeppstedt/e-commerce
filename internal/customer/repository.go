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
	getUserByID(userID int) (*user, error)
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
		"SELECT user_id, is_admin FROM users WHERE (username = ? OR email = ?) AND password = ?",
		loginInput, loginInput, password).
		Scan(&userID, &is_admin)
	fmt.Println("UserID:", userID)
	fmt.Println("is_admin:", is_admin)
	if err != nil {
		//Hanter Err.NoRows
		fmt.Println("ERROR userLogin in repo:", err)
	}
	return userID, is_admin, err

}

func (r *UserRepository) getUserByID(userID int) (*user, error) {
	var u user
	err := r.db.QueryRow(`
        SELECT username, email, first_name, last_name, address, zip_code, phone_number
        FROM users WHERE user_id = ?`, userID,
	).Scan(&u.UserName, &u.Email, &u.FirstName, &u.LastName, &u.Address, &u.ZipCode, &u.PhoneNumber)
	if err != nil {
		fmt.Println("Hur hamnade vi här? error in getUserByID repo:", err)
		return nil, err
	}
	return &u, nil
}
