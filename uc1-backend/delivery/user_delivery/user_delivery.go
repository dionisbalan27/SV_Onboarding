package user_delivery

import (
	helpers "backend-api/helpers/helpers_user"
	"backend-api/models/user/dto"
	"backend-api/usecase/user_usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserDelivery interface {
	GetAllUsers(*gin.Context)
	GetUserById(*gin.Context)
	CreateNewUser(*gin.Context)
	UpdateUserData(*gin.Context)
	DeleteUserById(*gin.Context)
	UserLogin(*gin.Context)
}

type userDelivery struct {
	usecase user_usecase.UserUsecase
}

func GetUserDelivery(userUsecase user_usecase.UserUsecase) UserDelivery {
	return &userDelivery{
		usecase: userUsecase,
	}
}

func (res *userDelivery) GetAllUsers(c *gin.Context) {
	response := res.usecase.GetAllUsers()
	// fmt.Printf("%+v", response)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) GetUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.GetUserById(id)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) CreateNewUser(c *gin.Context) {
	request := dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Error Bad Request", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}
	}
	response := res.usecase.CreateNewUser(request)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (res *userDelivery) UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	request := dto.User{}
	c.ShouldBindJSON(&request)
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Error Bad Request", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}

	}
	response := res.usecase.UpdateUserData(request, id)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteUserById(id)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) UserLogin(c *gin.Context) {
	request := dto.Login{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Error Bad Request", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}
	}
	response := res.usecase.UserLogin(request)
	if response.Status != "ok" {
		errorRes := helpers.ResponseError("Internal server error", 500)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.JSON(http.StatusOK, response)

}
