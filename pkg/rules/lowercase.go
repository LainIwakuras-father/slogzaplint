package rules

import (
	"errors"
	"unicode"
)

func Lowercase(str string) error {
	if str == "" {
		return nil
	}
	first := []rune(str)[0]
	if unicode.IsLetter(first) && unicode.IsLower(first) {
		return nil
	}
	return errors.New(msgLowercase)
}
