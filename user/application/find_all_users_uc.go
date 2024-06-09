package application

import userdomain "social_media/user/domain"

func FindAllUsersUC(rep userdomain.UserRepository) ([]string, error) {
	var users []string
	users, err := rep.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
