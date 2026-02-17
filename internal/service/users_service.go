package service

import (
	"ecommerce/duckyarmy/internal/models"
	"ecommerce/duckyarmy/internal/repository"
)

type UsersService interface {
	GetUsersByUserID(userID int) (models.User, error)
}

type usersService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UsersService {
	return &usersService{repo: r}
}

func (s *usersService) GetUsersByUserID(userID int) (models.User, error) {
	return s.repo.GetByUserID(userID)
}
