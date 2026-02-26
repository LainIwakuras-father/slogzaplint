package rules

import (
	"errors"
	"unicode"
)

func NoSpecial(str string) error {
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || IsPunctuationMarks(r) {
			continue
		}
		return errors.New(msgSpecial)
	}
	return nil
}

func IsPunctuationMarks(r rune) bool {
	return (r == ' ' || r == '_' || r == '-' || r == '=' || r == ':' || r == '.' || r == ',')
}
