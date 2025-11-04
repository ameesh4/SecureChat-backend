package service

import (
	"errors"
	"log"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/db/schema"
	"securechat/backend/src/utils"
)

type AuthService struct {
	jwtService *JWTService
}

func NewAuthService(jwtService *JWTService) *AuthService {
	return &AuthService{
		jwtService: jwtService,
	}
}

func RegisterUser(user *schema.User) (*schema.User, error) {
	if user == nil {
		return nil, errors.New("user cannot be nil")
	}

	if user.Email != "" {
		existingUser, _ := repository.GetUserByEmail(user.Email)
		if existingUser != nil {
			return nil, errors.New("user with this email already exists")
		}
	}

	if user.Email == "" {
		return nil, errors.New("email must be provided")
	}
	if user.Email != "" && !utils.ValidEmail(user.Email) {
		return nil, errors.New("invalid email format")
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

func (a *AuthService) AuthenticateUser(user model.LoginRequest) (*model.LoginResponse, error) {
	if user.Email == "" && user.Password == "" {
		return nil, errors.New("email and password cannot be empty")
	}

	var existingUser *schema.User

	if utils.ValidEmail(user.Email) {
		var err error
		existingUser, err = repository.GetUserByEmail(user.Email)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(user.Password, existingUser.Password) {
		return nil, errors.New("invalid password")
	}

	response := &model.LoginResponse{}
	accessToken, err := a.jwtService.GenerateToken(existingUser.Id)
	if err != nil {
		return nil, err
	}

	response.Token = accessToken
	if existingUser.RefreshToken != "" {
		res, err := a.jwtService.ValidateRefreshToken(existingUser.RefreshToken)
		if err != nil {
			log.Printf("Error validating refresh token: %v", err)
		}
		if res == nil {
			refreshToken, err := a.jwtService.GenerateRefreshToken(existingUser.Id)
			if err != nil {
				return nil, err
			}
			existingUser.RefreshToken = refreshToken
		}
	} else {
		refreshToken, err := a.jwtService.GenerateRefreshToken(existingUser.Id)
		if err != nil {
			return nil, err
		}
		existingUser.RefreshToken = refreshToken
	}
	_, err = repository.UpdateUser(existingUser)
	if err != nil {
		return nil, err
	}
	existingUser.Password = ""
	existingUser.RefreshToken = ""
	response.User = *existingUser
	return response, nil
}
