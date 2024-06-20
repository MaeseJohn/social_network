package application

import "social_media/user/domain"

func FindAllUsersUC(rep domain.UserRepository) ([]string, error) {
	var users []string
	users, err := rep.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
