package user_usecase

import (
	"backend-api/helpers"
	"backend-api/models/user/dto"
	"backend-api/models/user/entity"
	"backend-api/repository/user_repository"
	"backend-api/usecase/jwt_usecase"
	"errors"

	"gorm.io/gorm"
)

type UserUsecase interface {
	GetAllUsers() dto.Response
	GetUserById(string) dto.Response
	CreateNewUser(dto.User) dto.Response
	UpdateUserData(dto.User, string) dto.Response
	DeleteUserById(string) dto.Response
	UserLogin(dto.Login) dto.Response
}

type userUsecase struct {
	userRepo   user_repository.UserRepository
	jwtUsecase jwt_usecase.JwtUsecase
}
type UserUsecaseTest struct {
	userRepo *user_repository.UserRepositoryMock
}

func GetUserUsecase(jwtUsecase jwt_usecase.JwtUsecase, userRepository user_repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo:   userRepository,
		jwtUsecase: jwtUsecase,
	}
}

func (user *userUsecase) GetAllUsers() dto.Response {
	userlist, err := user.userRepo.GetAllUsers()
	response := []dto.UserList{}
	for _, user := range userlist {
		role := dto.Role{Id: user.RoleID, Title: user.Title}
		responseData := dto.UserList{
			Id:     user.ID,
			Name:   user.Name,
			Role:   role,
			Active: user.Active,
		}
		response = append(response, responseData)
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	return helpers.ResponseSuccess("ok", nil, response, 200)
}

func (user *userUsecase) GetUserById(id string) dto.Response {
	userData, err := user.userRepo.GetUserById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err.Error(), 500)
	}

	role := dto.Role{
		Id:    userData.RoleID,
		Title: userData.Title,
	}

	userResponse := map[string]interface{}{
		"id":             userData.ID,
		"name":           userData.Name,
		"email":          userData.Email,
		"role":           role,
		"personalNumber": userData.Personal_number,
		"active":         userData.Active,
	}
	return helpers.ResponseSuccess("ok", nil, userResponse, 200)
}

func (user *userUsecase) CreateNewUser(newUser dto.User) dto.Response {
	paswordHash, _ := helpers.HashPassword(newUser.Password)
	userInsert := entity.User{
		ID:              newUser.Id,
		Name:            newUser.Name,
		Email:           newUser.Email,
		Personal_number: newUser.Personal_number,
		Password:        paswordHash,
	}

	userData, err := user.userRepo.CreateNewUser(userInsert)

	if err != nil {
		if err.Error() == "Personal number already registered" {
			return helpers.ResponseError("Conflict", err.Error(), 409)
		} else {
			return helpers.ResponseError("Internal server error", err.Error(), 500)
		}

	}
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID}, 201)
}

func (user *userUsecase) UpdateUserData(userUpdate dto.User, id string) dto.Response {

	paswordHash, _ := helpers.HashPassword(userUpdate.Password)
	userInsert := entity.User{
		// ID: userUpdate.Id,
		Personal_number: userUpdate.Personal_number,
		Name:            userUpdate.Name,
		Password:        paswordHash,
		Email:           userUpdate.Email,
		Active:          userUpdate.Active,
		RoleID:          userUpdate.Role.Id,
	}

	_, err := user.userRepo.UpdateUserData(userInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		if err.Error() == "Personal number already taken" {
			return helpers.ResponseError("Confilct", err.Error(), 409)
		}
		return helpers.ResponseError("Internal server error", err.Error(), 500)
	}
	userUpdate.Id = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (user *userUsecase) DeleteUserById(id string) dto.Response {

	err := user.userRepo.DeleteUserById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err.Error(), 500)
	}
	return helpers.ResponseSuccess("ok", nil, nil, 200)
}

func (user *userUsecase) UserLogin(newUser dto.Login) dto.Response {

	res, role, err := user.userRepo.CheckLogin(newUser)
	if err != nil {
		return helpers.ResponseError("Data not found", "Wrong personal number / password", 404)
	}

	t, err := user.jwtUsecase.GenerateToken(res.ID)
	if err != nil {
		return helpers.ResponseError("Data not found", "Wrong personal number / password", 404)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"token": t, "name": res.Name, "role": role}, 200)

}
