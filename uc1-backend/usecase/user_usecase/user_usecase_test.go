package user_usecase

import (
	"backend-api/models/user/dto"
	"backend-api/models/user/entity"
	"backend-api/repository/user_repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = user_repository.UserRepositoryMock{Mock: mock.Mock{}}
var userUc = UserUsecaseTest{userRepo: &userRepository}

func TestGetAllUsersSuccess(t *testing.T) {
	expected := []entity.User{}
	userRepository.Mock.On("GetAllUsers").Return(expected, nil)
	result, err := userUc.userRepo.GetAllUsers()

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestGetUserByIdNotFound(t *testing.T) {
	userRepository.Mock.On("GetUserById", "1").Return(nil, errors.New("Data not found"))

	result, err := userUc.userRepo.GetUserById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestGetUserByIdSuccess(t *testing.T) {
	expected := &entity.User{}

	userRepository.Mock.On("GetUserById", "2").Return(expected, nil)

	result, err := userUc.userRepo.GetUserById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestUpdateUserNotFound(t *testing.T) {
	userRepository.Mock.On("GetUserById", "1").Return(nil, errors.New("Data not found"))

	result, err := userUc.userRepo.GetUserById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestUpdateUserSuccess(t *testing.T) {
	expected := &entity.User{}

	userRepository.Mock.On("UpdateUserData", "2").Return(expected, nil)

	result, err := userUc.userRepo.UpdateUserData("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewUserSuccess(t *testing.T) {
	request := dto.User{
		Name:            "New User",
		Personal_number: "123456789",
	}

	expected := &entity.User{}

	userRepository.Mock.On("CreateNewUser").Return(expected, nil)

	result, err := userUc.userRepo.CreateNewUser(request)

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewUserFailed(t *testing.T) {
	request := dto.User{
		Name:            "New User",
		Personal_number: "123456789",
	}

	expected := &entity.User{}

	userRepository.Mock.On("CreateNewUser").Return(nil, errors.New("Failed create User"))

	result, err := userUc.userRepo.CreateNewUser(request)

	assert.NotEqual(t, expected, result)
	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestDeleteUserFailed(t *testing.T) {
	userRepository.Mock.On("DeleteUserById", "1").Return(nil, errors.New("Data not found"))

	result, err := userUc.userRepo.DeleteUserById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestDeleteUserSuccess(t *testing.T) {
	expected := &entity.User{}

	userRepository.Mock.On("DeleteUserById", "2").Return(expected, nil)

	result, err := userUc.userRepo.DeleteUserById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}
