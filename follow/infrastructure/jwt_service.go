package infrastructure

import (
	"fmt"
	"social_media/follow/domain"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServiceImpl struct {
	SecretKey string
}

func NewJWTService(SecretKey string) *JwtServiceImpl {
	return &JwtServiceImpl{SecretKey: SecretKey}
}

func (service *JwtServiceImpl) ObtainTokenClaims(reqToken string) (jwt.MapClaims, error) {
	splitToken := strings.Split(reqToken, "Bearer ")
	tokenstring := splitToken[1]

	// Check err handling
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(service.SecretKey), nil
	})

	if err != nil {
		return nil, domain.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.ErrUnauthorized
	}
	return claims, nil

}
