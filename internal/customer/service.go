package customer

type userService interface {
	getUsersByUserID(userID int) (user, error)
}

// uhm better name maybe 😅
type userService1 struct {
	repo userRepository
}

func NewUserService1(r userRepository) *userService1 {
	return &userService1{repo: r}
}

func (s *userService1) getUsersByUserID(userID int) (user, error) {
	return s.repo.getByUserID(userID)
}
