package input

import (
	"errors"
)

type UserRegisterDto struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
}

func (dto *UserRegisterDto) Valid() error {
	if !emailRegex.MatchString(dto.Email) {
		return errors.New("invalid email")
	}

	if dto.Password == "" {
		return errors.New("invalid password")
	}

	if dto.Password != dto.RepeatPassword {
		return errors.New("invalid repeat password")
	}

	if dto.Lastname == "" {
		return errors.New("empty lastname")
	}

	if dto.Firstname == "" {
		return errors.New("empty firstname")
	}

	return nil
}
