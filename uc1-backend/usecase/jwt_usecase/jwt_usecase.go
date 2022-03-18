package jwt_usecase

import (
	"backend-api/repository/user_repository"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtUsecase interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	ValidateTokenAndGetUserId(string) (string, error)
}

type jwtUsecase struct {
	userRepo user_repository.UserRepository
}

func GetJwtUsecase(userRepository user_repository.UserRepository) JwtUsecase {
	return &jwtUsecase{
		userRepo: userRepository,
	}
}

type CustomClaim struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func (usecase *jwtUsecase) GenerateToken(userId string) (string, error) {
	claim := CustomClaim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
	// 	UserID: userId,
	// 	StandardClaims: jwt.StandardClaims{
	// 		IssuedAt:  time.Now().Unix(),
	// 		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
	// 		Issuer:    os.Getenv("APP_NAME"),
	// 	},
	// })
	// return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (usecase *jwtUsecase) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// claims := &CustomClaim{}
	// jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("SECRET_KEY")), nil
	// })
	// return claims.UserID
}

func (usecase *jwtUsecase) ValidateTokenAndGetUserId(token string) (string, error) {

	validatedToken, _ := usecase.ValidateToken(token)
	// if err != nil {
	// 	return "", err
	// }

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to claim token")
	}

	return claims["user_id"].(string), nil
	//return validatedToken, nil
}
