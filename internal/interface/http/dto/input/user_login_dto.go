package input

import (
	"errors"
	"regexp"
)

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (dto *UserLoginDto) Valid() error {
	if !emailRegex.MatchString(dto.Email) {
		return errors.New("invalid email")
	}

	if dto.Password == "" {
		return errors.New("invalid password")
	}

	return nil
}
