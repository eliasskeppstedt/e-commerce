package customer

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type userService interface {
	getUserByUsername(username string) (user, error)
	registerUser(username, password string) error
}

// uhm better name maybe 😅
type userService1 struct {
	repo userRepository
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewUserService1(r userRepository) *userService1 {
	return &userService1{repo: r}
}

func (s *userService1) getUserByUsername(username string) (user, error) {
	return s.repo.getUserByUsername(username)
}

func (s *userService1) registerUser(username, password, email string) error {
	return s.repo.registerUser(username, password, email)
}

func (s *userService1) userLogin(loginInput, password string) (string, error) {
	user, err := s.repo.userLogin(loginInput, password)
	token := ""
	if err != nil {
		return token, err
	}

	token, err = generateToken(uint(user.UserID))

	return token, err
}

func generateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	var jwtKey = []byte("your_secret_key")
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
