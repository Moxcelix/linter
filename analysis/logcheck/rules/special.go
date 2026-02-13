package rules

import (
	"errors"
	"slices"
)

var specialChars = []rune{'!', '?', '.', ',', ':', ';', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '=', '{', '}', '[', ']', '|', '\\', '/', '<', '>', '~', '`'}

var SpecialRuleError = errors.New("log message should not contain special characters")

func CheckSpecialRule(msg string) error {
	for _, char := range msg {
		if slices.Contains(specialChars, char) {
			return SpecialRuleError
		}
	}

	return nil
}
