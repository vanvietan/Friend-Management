package pkg

import (
	"errors"
	"net/mail"
)

// CheckValidEmail used for checking whether an email address parameter is valid or not
func CheckValidEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email address")
	}
	return nil
}
