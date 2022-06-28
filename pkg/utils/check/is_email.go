package check

import (
	"net/mail"
	"strings"
)

func IsEmail(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}

	if !strings.Contains(email, "@mercadolibre.com") {
		return false
	}
	return true

}
