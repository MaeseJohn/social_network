package infrastructure

import (
	"social_media/user/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServiceImpl struct {
	SecretKey string
}

func NewJWTService(SecretKey string) *JwtServiceImpl {
	return &JwtServiceImpl{SecretKey: SecretKey}
}

func (service *JwtServiceImpl) CreateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.SecretKey))
}
