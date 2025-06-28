package utils

import "regexp"

func ValidEmail(email string) bool {
	matched, err := regexp.Match(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/`, []byte(email))
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
