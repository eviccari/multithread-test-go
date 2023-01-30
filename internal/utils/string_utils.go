package utils

import "strings"

func IsEmptyString(s string) bool {
	return s == "" || s == " " || strings.TrimSpace(s) == ""
}
