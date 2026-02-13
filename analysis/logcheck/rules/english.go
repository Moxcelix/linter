package rules

import (
	"errors"
)

var EnglishRuleError = errors.New("log message should have english characters only")

func CheckEnglishRule(msg string) error {
	for _, r := range msg {
		if r > 127 {
			return EnglishRuleError
		}
	}

	return nil
}
