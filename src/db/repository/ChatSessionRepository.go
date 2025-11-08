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

	// Subquery to get the latest message ID for each session
	// Matches messages to sessions by participant IDs
	subquery := db.DB.Raw(`
		SELECT 
			cs.id as session_id,
			(SELECT cm.id 
			 FROM chat_messages cm
			 WHERE (
				(cs.participant1 = cm.sender_id AND cs.participant2 = cm.receiver_id) OR
				(cs.participant1 = cm.receiver_id AND cs.participant2 = cm.sender_id)
			 )
			 ORDER BY cm.created_at DESC
			 LIMIT 1
			) as latest_message_id
		FROM chat_sessions cs
	`)

	// Join sessions with the latest message
	result := db.DB.Preload("User1").Preload("User2").
		Table("chat_sessions").
		Select("chat_sessions.*").
		Joins("LEFT JOIN (?) as session_latest ON chat_sessions.id = session_latest.session_id", subquery).
		Joins("LEFT JOIN chat_messages ON chat_messages.id = session_latest.latest_message_id").
		Where("chat_sessions.participant1 = ? OR chat_sessions.participant2 = ?", userId, userId).
		Order("chat_sessions.updated_at DESC").
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
