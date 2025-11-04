package service

import (
	"securechat/backend/src/db/repository"
)

type AllUserResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func GetAllUsers() ([]AllUserResponse, error) {
	var filteredUsers []AllUserResponse
	users, err := repository.GetAllUsers()
	for _, user := range users {
		filteredUsers = append(filteredUsers, AllUserResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
	if err != nil {
		return nil, err
	}
	return filteredUsers, nil
}
