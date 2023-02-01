package mockinggo

import "github.com/stretchr/testify/mock"

type UserRepoMock struct {
	Mock mock.Mock
}

func (r *UserRepoMock) FindAll() *[]User {
	args := r.Mock.Called()
	if args.Get(0) == nil {
		return nil
	}

	users := args.Get(0).([]User)
	return &users
}
