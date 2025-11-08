package repository

import (
	"securechat/backend/src/db"
	"securechat/backend/src/db/schema"
	"securechat/backend/src/utils"
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
	var user1 *schema.User
	var user2 *schema.User
	if userId1 == session.User1.Id {
		user1 = &session.User1
		user2 = &session.User2
	} else {
		user1 = &session.User2
		user2 = &session.User1
	}
	returningSession := &schema.ChatSession{
		Id:           session.Id,
		Participant1: userId1,
		Participant2: userId2,
		A1:           utils.Ternery(userId1 == session.Participant1, session.A1, session.A2),
		A2:           utils.Ternery(userId1 == session.Participant1, session.A2, session.A1),
		CreatedAt:    session.CreatedAt,
		UpdatedAt:    session.UpdatedAt,
		User1:        *user1,
		User2:        *user2,
	}

	return returningSession, nil
}

func GetChatSessionsByUserID(userId uint) ([]schema.ChatSession, error) {
	var sessions []schema.ChatSession

	result := db.DB.Preload("User1", utils.GeneralizeUser).Preload("User2", utils.GeneralizeUser).
		Where("participant1 = ? OR participant2 = ?", userId, userId).
		Order("updated_at DESC").
		Find(&sessions)

	if result.Error != nil {
		return nil, result.Error
	}

	return utils.Ternery(len(sessions) > 0, utils.GeneralizeSession(sessions, userId), []schema.ChatSession{}), nil
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
