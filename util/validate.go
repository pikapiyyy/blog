package util

import (
	"fmt"
	"regexp"
)

func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

func CheckUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}

func CheckPassword(password string) string {
	return checkLen("Password", password, 6, 50)
}

func CheckEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}
