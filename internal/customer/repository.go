package customer

import (
	"database/sql"
	"errors"
	"fmt"
)

type userRepository interface {
	getUserByUsername(username string) (user, error)
	create(username, password string) error
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) *mysqlUserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) getUserByUsername(username string) (user, error) {

	query := "SELECT userID, userName, password FROM users WHERE userID = ?"

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

func (r *mysqlUserRepository) create(username, password string) error {
	var existingUsername string

	if r.db == nil {
		fmt.Println(r.db)
		return errors.New("DB conenction doesnt exist in repo")
	}

	err := r.db.QueryRow(
		"SELECT userName FROM users WHERE userName = ?",
		username,
	).Scan(&existingUsername)

	if username == existingUsername {
		return errors.Join(err, errors.New("^joined error: User already exists"))
	}

	dummyAddress := username + "." + password + "@gmail.com"
	_, err = r.db.Exec(
		"INSERT INTO users (userName, password, emailAddress) VALUES (?, ?, ?)",
		username, password, dummyAddress,
	)

	if err != nil {
		return errors.Join(err, errors.New("^joined error: Cant create user"))
	}

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

	return nil
}
