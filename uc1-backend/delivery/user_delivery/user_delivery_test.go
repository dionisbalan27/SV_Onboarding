package user_delivery

import (
	"backend-api/models/user/dto"
	"backend-api/usecase/user_usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userUsecase = user_usecase.UserUsecaseMock{Mock: mock.Mock{}}
var userDel = UserDeliveryTest{userUsecase: &userUsecase}

func TestGetAllUsersSuccess(t *testing.T) {

	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       []dto.UserList{},
	}

	userUsecase.Mock.On("GetAllUsers").Return(expected)

	result := userDel.userUsecase.GetAllUsers()

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}

func TestGetProductByIdSuccess(t *testing.T) {

	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       dto.User{},
	}

	userUsecase.Mock.On("GetUserById", "1").Return(expected)

	result := userDel.userUsecase.GetUserById("1")

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}

func TestGetProductByIdNotFound(t *testing.T) {
	expected := dto.Response{
		StatusCode: 404,
		Status:     "Data not found",
		Error:      errors.New("Record not found"),
		Data:       nil,
	}

	userUsecase.Mock.On("GetUserById", "2").Return(expected)

	result := userDel.userUsecase.GetUserById("2")

	assert.Equal(t, expected.Status, result.Status)
	assert.Nil(t, result.Data)
	assert.NotNil(t, result.Error)
}

func TestDeleteProductSuccess(t *testing.T) {
	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       map[string]interface{}{"id": "1"},
	}

	userUsecase.Mock.On("DeleteUserById", "1").Return(expected)

	result := userDel.userUsecase.DeleteUserById("1")

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestDeleteProductFailed(t *testing.T) {

	expected := dto.Response{
		StatusCode: 404,
		Status:     "Data not found",
		Error:      errors.New("Record not found"),
		Data:       nil,
	}

	userUsecase.Mock.On("DeleteUserById", "2").Return(expected)

	result := userDel.userUsecase.DeleteUserById("2")

	assert.Equal(t, expected.Error, result.Error)
	assert.Nil(t, result.Data)
	assert.NotNil(t, result.Error)
}

func TestUpdateUserDataSuccess(t *testing.T) {
	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       map[string]interface{}{"id": "1"},
	}

	userUsecase.Mock.On("UpdateUserData", "1").Return(expected)

	result := userDel.userUsecase.UpdateUserData("1")

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestUpdateUserDataFailed(t *testing.T) {

	expected := dto.Response{
		StatusCode: 404,
		Status:     "Data not found",
		Error:      errors.New("Record not found"),
		Data:       nil,
	}

	userUsecase.Mock.On("UpdateUserData", "2").Return(expected)

	result := userDel.userUsecase.UpdateUserData("2")

	assert.Equal(t, expected.Error, result.Error)
	assert.Nil(t, result.Data)
	assert.NotNil(t, result.Error)
}

func TestCreateNewUserSuccess(t *testing.T) {
	userData := dto.User{}
	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       map[string]interface{}{"id": "1"},
	}

	userUsecase.Mock.On("CreateNewUser").Return(expected)

	result := userDel.userUsecase.CreateNewUser(userData)

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestCreateNewUserConflict(t *testing.T) {
	userData := dto.User{Personal_number: "1"}
	expected := dto.Response{
		StatusCode: 409,
		Status:     "Conflict",
		Error:      errors.New("Personal number already exist"),
		Data:       nil,
	}

	userUsecase.Mock.On("CreateNewUser").Return(expected)

	result := userDel.userUsecase.CreateNewUser(userData)

	assert.Nil(t, result.Data)
	assert.NotNil(t, result.Error)
}
