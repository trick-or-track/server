package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Format(str string) string {
	return strings.Trim(strings.ToLower(str), " ")
}

func Proper(str string) string {
	c := cases.Title(language.English)
	return c.String(str)
}
