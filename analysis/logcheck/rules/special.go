package rules

import (
	"errors"
	"slices"
)

var SpecialRuleError = errors.New("log message should not contain special characters")

type SpecialRule struct {
	specialChars []rune
}

type SpecialSymbolsProvider interface {
	Provide() []rune
}

func NewSpecialRule(specialSymbolsProvider SpecialSymbolsProvider) *SpecialRule {
	return &SpecialRule{
		specialChars: specialSymbolsProvider.Provide(),
	}
}

func (rule *SpecialRule) Check(msg string) error {
	for _, char := range msg {
		if slices.Contains(rule.specialChars, char) {
			return SpecialRuleError
		}
	}

	return nil
}
