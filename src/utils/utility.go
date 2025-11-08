package utils

import (
	"regexp"

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

func FindKeysByValueConnections(m map[uint]*socket.Socket, targetValue *socket.Socket) []uint {
	var keys []uint
	for k, v := range m {
		if v == targetValue {
			keys = append(keys, k)
		}
	}
	return keys
}

func Ternery[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
