package user_delivery

import (
	"backend-api/helpers"
	"backend-api/models/user/dto"
	"backend-api/usecase/user_usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
type UserDeliveryTest struct {
	userUsecase *user_usecase.UserUsecaseMock
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
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) GetUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.GetUserById(id)
	if response.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusOK, response)
		return
	}

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) CreateNewUser(c *gin.Context) {
	request := dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := res.usecase.CreateNewUser(request)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (res *userDelivery) UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	request := dto.User{}
	c.ShouldBindJSON(&request)

	response := res.usecase.UpdateUserData(request, id)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteUserById(id)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) UserLogin(c *gin.Context) {
	request := dto.Login{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400)
		fmt.Printf("%+v", errorRes)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := res.usecase.UserLogin(request)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)

}
