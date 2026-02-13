package rules

import (
	"errors"
	"unicode"
)

var LowercaseRuleError = errors.New("log message should start with a lowercase letter")

func CheckLowercaseRule(msg string) error {
	if len(msg) == 0 {
		return nil
	}

	firstChar := []rune(msg)[0]
	if unicode.IsUpper(firstChar) {
		return LowercaseRuleError
	}

	return nil
}
