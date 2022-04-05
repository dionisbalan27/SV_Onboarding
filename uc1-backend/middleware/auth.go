package middleware

import (
	"net/http"

	"backend-api/helpers"
	"backend-api/repository/user_repository"
	"backend-api/usecase/jwt_usecase"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader[7:])
		if err != nil {
			errorRes := helpers.ResponseError("Unauthorized", 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		c.Set("user_id", userId)
	}
}

func JWTAuthAdmin(jwtUsecase jwt_usecase.JwtUsecase, userRepo user_repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader[7:])
		if err != nil {
			errorRes := helpers.ResponseError("Unauthorized", 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		user, err := userRepo.GetUserById(userId)
		if err != nil {
			errorRes := helpers.ResponseError("Internal Server Error", 500)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorRes)
			return
		}

		if user.Title != "admin" {
			errorRes := helpers.ResponseError("Unauthorized", 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}
		// panic(user.Title)
		c.Set("user_id", userId)
	}
}

func JWTAuthChecker(jwtUsecase jwt_usecase.JwtUsecase, userRepo user_repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader[7:])
		if err != nil {
			errorRes := helpers.ResponseError("Unauthorized", 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		user, err := userRepo.GetUserById(userId)
		if err != nil {
			errorRes := helpers.ResponseError("Internal Server Error", 500)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorRes)
			return
		}

		if user.Title != "checker" {
			if user.Title != "admin" {
				errorRes := helpers.ResponseError("Unauthorized", 401)
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
				return
			}
			c.Set("user_id", userId)
		}

		// panic(user.Title)
		c.Set("user_id", userId)
	}
}

func JWTAuthSigner(jwtUsecase jwt_usecase.JwtUsecase, userRepo user_repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader[7:])
		if err != nil {
			errorRes := helpers.ResponseError("Unauthorized", 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		user, err := userRepo.GetUserById(userId)
		if err != nil {
			errorRes := helpers.ResponseError("Internal Server Error", 500)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorRes)
			return
		}

		if user.Title != "signer" {
			if user.Title != "admin" {
				errorRes := helpers.ResponseError("Unauthorized", 401)
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
				return
			}
			c.Set("user_id", userId)
		}
		// panic(user.Title)
		c.Set("user_id", userId)
	}
}
