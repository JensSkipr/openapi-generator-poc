/* This file is auto-generated, manual edits in this file will be overwritten! */
package utils

import (
	"net/mail"
	"strings"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}
