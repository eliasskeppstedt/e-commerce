package customer

type userService interface {
	getUserByUsername(username string) (user, error)
	registerUser(username, password string) error
}

// uhm better name maybe 😅
type userService1 struct {
	repo userRepository
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
