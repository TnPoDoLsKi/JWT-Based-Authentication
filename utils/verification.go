package utils

import (
	"regexp"
)

func EmailValidation(email string) (bool, error) {
	reg := `^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$`

	return regexp.MatchString(reg, email)
}

func PasswordValidation(password string) (bool, error) {
	reg := `^(.{0,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{4,})|(.{1,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{3,})|(.{2,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{2,})|(.{3,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{1,})|(.{4,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{0,})$
`

	return regexp.MatchString(reg, password)
}
