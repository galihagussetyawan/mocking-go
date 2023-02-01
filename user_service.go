package mockinggo

import "errors"

type UserService struct {
	UserRepo UserRepository
}

func (s *UserService) FindAll() (*[]User, error) {
	users := s.UserRepo.FindAll()
	if users == nil {
		return users, errors.New("users is empty")
	} else {
		return users, nil
	}
}
