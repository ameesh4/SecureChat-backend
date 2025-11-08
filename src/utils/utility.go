package utils

import (
	"errors"
	"regexp"
	"securechat/backend/src/db/schema"

	"github.com/zishang520/socket.io/v2/socket"
)

func ValidEmail(email string) bool {
	matched, err := regexp.Match(`^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$`, []byte(email))
	if err != nil {
		return false
	}
	return matched
}

func ValidPhoneNumber(phone string) bool {
	matched, err := regexp.Match(`^\+?[1-9]\d{1,14}$`, []byte(phone))
	if err != nil {
		return false
	}
	return matched
}

func FindKeysByValueConnections(m map[uint]*socket.Socket, targetValue *socket.Socket) ([]uint, error) {
	var keys []uint
	for k, v := range m {
		if v == targetValue {
			keys = append(keys, k)
		}
	}
	return keys, Ternery(len(keys) > 0, nil, errors.New("NOT FOUND"))
}

func Ternery[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func GeneralizeSession(sessions []schema.ChatSession, userId1 uint) []schema.ChatSession {
	var generalizedSessions []schema.ChatSession
	for _, session := range sessions {
		var user1 *schema.User
		var user2 *schema.User
		if userId1 == session.User1.Id {
			user1 = &session.User1
			user2 = &session.User2
		} else {
			user1 = &session.User2
			user2 = &session.User1
		}
		generalizedSession := &schema.ChatSession{
			Id:           session.Id,
			Participant1: userId1,
			Participant2: Ternery(userId1 == session.Participant1, session.Participant2, session.Participant1),
			A1:           Ternery(userId1 == session.Participant1, session.A1, session.A2),
			A2:           Ternery(userId1 == session.Participant1, session.A2, session.A1),
			CreatedAt:    session.CreatedAt,
			UpdatedAt:    session.UpdatedAt,
			User1:        *user1,
			User2:        *user2,
		}
		generalizedSessions = append(generalizedSessions, *generalizedSession)
	}
	return generalizedSessions
}
