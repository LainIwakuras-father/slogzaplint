package rules

import (
	"errors"
	"strings"
)

func NoSensitive(str string, patterns []string) error {
	for _, pattern := range patterns {
		if strings.Contains(strings.ToLower(str), pattern) {
			return errors.New(msgSensitive)
		}
	}
	return nil
}
