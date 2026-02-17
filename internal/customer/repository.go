package customer

import (
	"database/sql"
)

type UserRepository interface {
	GetByUserID(userID int) (User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUserID(userID int) (User, error) {

	query := "SELECT userID, userName, password FROM users WHERE userID = ?"

	var user User

	err := r.db.QueryRow(query, userID).
		Scan(&user.UserID, &user.UserName, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, err
		}
		return User{}, err
	}

	return user, nil
}
