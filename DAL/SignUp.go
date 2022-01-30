package DAL

import (
	"cloudAuthentication/DAO"
	"cloudAuthentication/config"
)

func SignUp(input DAO.UserRequestModel) (DAO.Users, error) {
	temp := DAO.Users{
		Username: input.Username,
		Password: input.Password,
	}
	err := config.DB.Create(&temp)
	if err != nil {
		return DAO.Users{}, err.Error
	}
	return temp, nil
}

func UsernameExists(username string) error {
	var user DAO.Users
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
