package service

import (
	"errors"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/db/schema"
)

func SendMessage(message model.Message) (*schema.ChatMessage, error) {
	_, err := repository.GetUserByID(message.ReceiverId)
	if err != nil {
		return nil, errors.New("receiver not found")
	}
	chatMessage := &schema.ChatMessage{
		SenderId:   message.SenderId,
		SessionId:  message.SessionId,
		ReceiverId: message.ReceiverId,
		Content:    message.Content,
		Iv:         message.Iv,
	}
	savedMessage, err := repository.CreateChatMessage(chatMessage)
	if err != nil {
		return nil, errors.New("failed to send message")
	}
	return savedMessage, nil
}

func GetChatMessages(sessionId uint) ([]schema.ChatMessage, error) {
	_, err := repository.GetChatSessionByID(sessionId)
	if err != nil {
		return nil, errors.New("session not found")
	}
	messages, err := repository.GetChatMessagesBySessionID(sessionId, 0, 0)
	if err != nil {
		return nil, errors.New("failed to get messages")
	}
	return messages, nil
}
