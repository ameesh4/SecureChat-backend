package repository

import (
	"securechat/backend/src/db"
	"securechat/backend/src/db/schema"
)

func CreateChatSession(session *schema.ChatSession) (*schema.ChatSession, error) {
	result := db.DB.Preload("User1").Preload("User2").Create(session)
	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func GetChatSessionByID(id uint) (*schema.ChatSession, error) {
	var session schema.ChatSession
	result := db.DB.Preload("User1").Preload("User2").First(&session, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}

func GetChatSessionBetweenUsers(userId1, userId2 uint) (*schema.ChatSession, error) {
	var session schema.ChatSession
	result := db.DB.Preload("User1").Preload("User2").
		Where("(participant1 = ? AND participant2 = ?) OR (participant1 = ? AND participant2 = ?)",
			userId1, userId2, userId2, userId1).
		First(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}

func GetChatSessionsByUserID(userId uint) ([]schema.ChatSession, error) {
	var sessions []schema.ChatSession
	result := db.DB.Preload("User1").Preload("User2").
		Where("participant1 = ? OR participant2 = ?", userId, userId).
		Order("updated_at DESC").
		Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}

	return sessions, nil
}

func UpdateChatSession(session *schema.ChatSession) (*schema.ChatSession, error) {
	if session == nil {
		return nil, nil
	}
	result := db.DB.Save(session)
	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func DeleteChatSession(id uint) error {
	var session schema.ChatSession
	result := db.DB.First(&session, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.DB.Delete(&session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
