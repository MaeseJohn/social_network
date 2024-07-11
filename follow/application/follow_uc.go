package application

import (
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	ObtainTokenClaims(reqToken string) (jwt.MapClaims, error)
}

func FollowUC(reqToken, followedId string, rep domain.FollowRepository, jwt JWTService) error {
	claims, err := jwt.ObtainTokenClaims(reqToken)
	if err != nil {
		return domain.ErrUnauthorized
	}

	followerId := claims["user_id"].(string)
	follow := domain.NewFollow(followerId, followedId)
	err = rep.SaveFollow(follow)
	if err != nil {
		return err
	}

	return nil
}
