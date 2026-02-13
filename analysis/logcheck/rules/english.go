package rules

import (
	"errors"
)

var EnglishRuleError = errors.New("log message should have english characters only")

type EnglishRule struct{}

func NewEnglishRule() *EnglishRule {
	return &EnglishRule{}
}

func (rule *EnglishRule) Check(msg string) error {
	for _, r := range msg {
		if r > 127 {
			return EnglishRuleError
		}
	}

	return nil
}
