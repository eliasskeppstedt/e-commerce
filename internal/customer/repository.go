package customer

import (
	"database/sql"
)

type userRepository interface {
	getUserByUsername(username string) (user, error)
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
		Scan(&user.UserID, &user.UserName, &user.Password, &user.EmailAddress)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	return user, nil
}
