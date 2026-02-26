// checkE
package rules

import (
	"errors"
	"unicode"
)

func English(str string) error {
	for _, r := range str {
		if unicode.IsLetter(r) && !IsLatin(r) {
			return errors.New(msgEnglish)
		}
	}
	return nil
}

func IsLatin(r rune) bool {
	return (r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z')
}
