package customer

import (
	"database/sql"
)

type userRepository interface {
	getByUserID(userID int) (user, error)
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) *mysqlUserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) getByUserID(userID int) (user, error) {

	query := "SELECT userID, userName, password FROM users WHERE userID = ?"

	var user user

	err := r.db.QueryRow(query, userID).
		Scan(&user.UserID, &user.UserName, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	return user, nil
}
