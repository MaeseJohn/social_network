package application

import "social_media/user/domain"

func FindAllUsersUC(rep domain.UserRepository) ([]domain.User, error) {
	var users []domain.User
	users, err := rep.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
