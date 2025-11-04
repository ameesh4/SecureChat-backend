package repository

import (
	"securechat/backend/src/db"
	"securechat/backend/src/db/schema"
)

func CreateUser(user *schema.User) (*schema.User, error) {
	result := db.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByID(id uint) (*schema.User, error) {
	var user schema.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func UpdateUser(user *schema.User) (*schema.User, error) {
	if user == nil {
		return nil, nil
	}
	result := db.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func DeleteUser(id uint) error {
	var user schema.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllUsers() ([]schema.User, error) {
	var users []schema.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []schema.User{}, nil
	}
	return users, nil
}
