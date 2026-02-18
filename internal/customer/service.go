package customer

type userService interface {
	getByID(userID int) (user, error)
	register(username, password string) error
}

// uhm better name maybe 😅
type userService1 struct {
	repo userRepository
}

func NewUserService1(r userRepository) *userService1 {
	return &userService1{repo: r}
}

func (s *userService1) getByID(userID int) (user, error) {
	return s.repo.getByUserID(userID)
}

func (s *userService1) register(username, password string) error {
	// eventualla krav för hur un och pwd måste se ut
	return s.repo.create(username, password)
}
