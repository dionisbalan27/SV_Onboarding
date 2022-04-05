package user_usecase

import (
	"backend-api/helpers"
	"backend-api/models/user/dto"
	"backend-api/models/user/entity"
	"errors"

	"github.com/stretchr/testify/mock"
)

type userUsecaseMockInterface interface {
	GetAllUsers() dto.Response
	GetUserById(id string) dto.Response
	UpdateUserData(id string) dto.Response
	CreateNewUser(user entity.User) dto.Response
	DeleteUserById(id string) dto.Response
}

type UserUsecaseMock struct {
	Mock mock.Mock
}

func (user *UserUsecaseMock) GetAllUsers() dto.Response {
	arguments := user.Mock.Called()
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("No record"), 404)
	}
	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}

func (user *UserUsecaseMock) GetUserById(id string) dto.Response {
	arguments := user.Mock.Called(id)
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("No record"), 404)
	}
	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}

func (user *UserUsecaseMock) UpdateUserData(id string) dto.Response {
	arguments := user.Mock.Called(id)

	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("No record"), 404)
	}
	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}

func (user *UserUsecaseMock) DeleteUserById(id string) dto.Response {
	arguments := user.Mock.Called(id)
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Data not found", errors.New("No record"), 404)
	}
	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}

func (user *UserUsecaseMock) CreateNewUser(userData dto.User) dto.Response {
	id := "1"
	arguments := user.Mock.Called()
	if arguments.Get(0) == nil {
		return helpers.ResponseError("Internal Server Error", errors.New("Failed to create user"), 500)
	}

	if userData.Personal_number == id {
		return helpers.ResponseError("Conflict", errors.New("Personal number already registered"), 409)
	}

	response := arguments.Get(0).(dto.Response)
	return helpers.ResponseSuccess(response.Status, response.Error, response.Data, response.StatusCode)
}
