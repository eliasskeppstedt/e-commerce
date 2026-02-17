package repository

import (
	"database/sql"
	"ecommerce/duckyarmy/internal/models"
)

type UserRepository interface {
	GetByUserID(userID int) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUserID(userID int) (models.User, error) {

	query := "SELECT userID, userName, password FROM users WHERE userID = ?"

	var user models.User

	err := r.db.QueryRow(query, userID).
		Scan(&user.UserID, &user.UserName, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
		return models.User{}, err
	}

	return user, nil
}
