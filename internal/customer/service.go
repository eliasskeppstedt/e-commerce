package customer

import (
	"ecommerce/duckyarmy/internal/cart"
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
		phone_number string) (int, error)
	userLogin(loginInput, password string) (int, bool, error)
	getUserByID(userID int) (*user, error)
}

// uhm better name maybe 😅
type userService1 struct {
	uRepo userRepository
	cRepo cart.CartRepository
}

func NewUserService1(r userRepository, c cart.CartRepository) *userService1 {
	return &userService1{uRepo: r, cRepo: c}
}

func (s *userService1) registerUser(username,
	password,
	email,
	first_name,
	last_name,
	address,
	zip_code,
	phone_number string) (int, error) {

	return s.uRepo.registerUser(username,
		password,
		email,
		first_name,
		last_name,
		address,
		zip_code,
		phone_number)
}

func (s *userService1) userLogin(loginInput, password string) (int, bool, error) {
	userID, isAdmin, err := s.uRepo.userLogin(loginInput, password)
	if err != nil {
		fmt.Println("error userLogin in service:", err)
		return -1, false, err //-1 errorcode borde vi göra någonting med den
	}
	return userID, isAdmin, err

}

func (s *userService1) getUserByID(userID int) (*user, error) {
	return s.uRepo.getUserByID(userID)
}
