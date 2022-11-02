package pkg

import (
	"errors"
	"regexp"
	"strings"
)

// CheckValidEmail used for checking whether an email address parameter is valid or not
func CheckValidEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if strings.TrimSpace(email) == "" || !emailRegex.MatchString(email) {
		return errors.New("invalid email address")
	}

	return nil
}
