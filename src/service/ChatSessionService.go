package service

import (
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/db/schema"
)

func NewChatSessionService(user *schema.User, email string) (*model.PublicKeys, error) {
	user2, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return &model.PublicKeys{
		User1PublicKey: user.PublicKey,
		User2PublicKey: user2.PublicKey,
	}, nil
}

func CreateChatSession(user1 schema.User, request model.CreateSessionRequest) (*schema.ChatSession, error) {
	user2, err := repository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	existingSession, err := repository.GetChatSessionBetweenUsers(user1.Id, user2.Id)
	if err == nil && existingSession != nil {
		return existingSession, nil
	}

	session := &schema.ChatSession{
		Participant1: user1.Id,
		Participant2: user2.Id,
		A1:           request.A1,
		A2:           request.A2,
	}

	session, err = repository.CreateChatSession(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}
