package user_usecase

import (
	helpers "backend-api/helpers/helpers_user"
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
		return helpers.ResponseError("Error Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}
	return helpers.ResponseSuccess("ok", nil, response)
}

func (user *userUsecase) GetUserById(id string) dto.Response {
	userData, err := user.userRepo.GetUserById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", nil)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", nil)
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
	return helpers.ResponseSuccess("ok", nil, userResponse)
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
		return helpers.ResponseError("Internal server error", err)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID})
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
		return helpers.ResponseError("Error Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	userUpdate.Id = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id})
}

func (user *userUsecase) DeleteUserById(id string) dto.Response {

	err := user.userRepo.DeleteUserById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Error Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("User deleted successfully", 200, nil)
}

func (user *userUsecase) UserLogin(newUser dto.Login) dto.Response {

	res, err := user.userRepo.CheckLogin(newUser)
	if err != nil {
		return helpers.ResponseError("Error Data not found", 404)
	}

	t, err := user.jwtUsecase.GenerateToken(res)
	if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"token": t})

}
