package mockinggo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &UserRepoMock{Mock: mock.Mock{}}
var userService = UserService{UserRepo: userRepository}

func TestUserService_FindAll(t *testing.T) {
	users := []User{
		{Name: "galih", Address: "trenggalek"},
	}

	userRepository.Mock.On("FindAll").Return(users)

	result := userService.UserRepo.FindAll()

	assert.NotNil(t, result)
	assert.Equal(t, len(*result), len(users))
}
