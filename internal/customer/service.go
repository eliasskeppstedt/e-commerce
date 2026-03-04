package customer

import (
	"fmt"
)

type userService interface {
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

// uhm better name maybe 😅
type userService1 struct {
	repo userRepository
}

func NewUserService1(r userRepository) *userService1 {
	return &userService1{repo: r}
}

func (s *userService1) registerUser(username,
	password,
	email,
	first_name,
	last_name,
	address,
	zip_code,
	phone_number string) error {

	return s.repo.registerUser(username,
		password,
		email,
		first_name,
		last_name,
		address,
		zip_code,
		phone_number)
}

func (s *userService1) userLogin(loginInput, password string) (int, bool, error) {
	userID, isAdmin, err := s.repo.userLogin(loginInput, password)
	if err != nil {
		fmt.Println("error userLogin in service:", err)
		return -1, false, err //-1 errorcode borde vi göra någonting med den
	}
	return userID, isAdmin, err

}

func (s *userService1) getUserByID(userID int) (*user, error) {
	return s.repo.getUserByID(userID)
}
