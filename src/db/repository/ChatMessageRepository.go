package repository

import (
	"securechat/backend/src/db"
	"securechat/backend/src/db/schema"
)

func CreateChatMessage(message *schema.ChatMessage) (*schema.ChatMessage, error) {
	result := db.DB.Create(message)
	if result.Error != nil {
		return nil, result.Error
	}

	return message, nil
}

func GetChatMessageByID(id uint) (*schema.ChatMessage, error) {
	var message schema.ChatMessage
	result := db.DB.Preload("Session").Preload("Sender").First(&message, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &message, nil
}

func GetChatMessagesBySessionID(sessionId uint, limit, offset int) ([]schema.ChatMessage, error) {
	var messages []schema.ChatMessage
	query := db.DB.Preload("Session").Preload("Sender").
		Where("session_id = ?", sessionId).
		Order("created_at ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	result := query.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func GetChatMessagesByUserID(userId uint, limit, offset int) ([]schema.ChatMessage, error) {
	var messages []schema.ChatMessage
	query := db.DB.Preload("Session").Preload("Sender").
		Where("sender_id = ?", userId).
		Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	result := query.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func UpdateChatMessage(message *schema.ChatMessage) (*schema.ChatMessage, error) {
	if message == nil {
		return nil, nil
	}
	result := db.DB.Save(message)
	if result.Error != nil {
		return nil, result.Error
	}

	return message, nil
}

func MarkMessageAsRead(id uint) error {
	result := db.DB.Model(&schema.ChatMessage{}).Where("id = ?", id).Update("is_read", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteChatMessage(id uint) error {
	var message schema.ChatMessage
	result := db.DB.First(&message, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.DB.Delete(&message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUnreadMessagesBySessionID(sessionId uint) ([]schema.ChatMessage, error) {
	var messages []schema.ChatMessage
	result := db.DB.Preload("Session").Preload("Sender").
		Where("session_id = ? AND is_read = ?", sessionId, false).
		Order("created_at ASC").
		Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}
