package service

import (
	"errors"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/db/schema"
	"securechat/backend/src/utils"
)

func RegisterUser(user *schema.User) (*schema.User, error) {
	if user == nil {
		return nil, errors.New("user cannot be nil")
	}

	if user.Email == "" || user.PhoneNumber == "" {
		return nil, errors.New("either email or phone number must be provided")
	}

	if user.Email != "" && !utils.ValidEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}

	if user.PhoneNumber != "" && !utils.ValidPhoneNumber(user.PhoneNumber) {
		return nil, errors.New("invalid phone number format")
	}

	if user.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	createdUser.Password = "" // Do not return password in response
	return createdUser, nil
}

func AuthenticateUser(emailOrPhone string, password string) (*model.LoginResponse, error) {
	if emailOrPhone == "" || password == "" {
		return nil, errors.New("email or phone number and password cannot be empty")
	}
	var user *schema.User

	if utils.ValidEmail(emailOrPhone) {
		var err error
		user, err = repository.GetUserByEmail(emailOrPhone)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	if utils.ValidPhoneNumber(emailOrPhone) {
		var err error
		user, err = repository.GetUserByPhoneNumber(emailOrPhone)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	user.Password = "" // Do not return password in response
	var response *model.LoginResponse
	response.User = *user

}
