package customer

type UsersService interface {
	GetUsersByUserID(userID int) (User, error)
}

type usersService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UsersService {
	return &usersService{repo: r}
}

func (s *usersService) GetUsersByUserID(userID int) (User, error) {
	return s.repo.GetByUserID(userID)
}
