package user_repository

import (
	"backend-api/models/product/entity"
	"backend-api/models/user/dto"
	"errors"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) GetAllUsers() ([]entity.User, error) {
	args := repository.Mock.Called()

	if args.Get(0) == nil {
		return nil, errors.New("Data not found")
	} else {
		user := args.Get(0).([]entity.User)
		return user, nil
	}
}

func (repository *UserRepositoryMock) GetUserById(id string) (*entity.User, error) {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, errors.New("Data not found")
	} else {
		user := args.Get(0).(*entity.User)
		return user, nil
	}
}

func (repository *UserRepositoryMock) UpdateUserData(id string) (*entity.User, error) {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, errors.New("Data not found")
	} else {
		user := args.Get(0).(*entity.User)
		return user, nil
	}
}

func (repository *UserRepositoryMock) DeleteUserById(id string) (*entity.User, error) {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, errors.New("Data not found")
	} else {
		user := args.Get(0).(*entity.User)
		return user, nil
	}
}

func (repository *UserRepositoryMock) CreateNewUser(newUser dto.User) (*entity.User, error) {
	args := repository.Mock.Called()

	if args.Get(0) == nil {
		return nil, errors.New("Failed create User")
	} else {

		user := args.Get(0).(*entity.User)

		return user, nil
	}
}
