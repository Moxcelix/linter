package rules

import (
	"errors"
	"slices"
)

var SpecialRuleError = errors.New("log message should not contain special characters")

type SpecialRule struct {
	specialChars []string
}

type SpecialSymbolsProvider interface {
	Provide() []string
}

func NewSpecialRule(specialSymbolsProvider SpecialSymbolsProvider) *SpecialRule {
	return &SpecialRule{
		specialChars: specialSymbolsProvider.Provide(),
	}
}

func (rule *SpecialRule) Check(msg string) error {
	for _, char := range msg {
		if slices.Contains(rule.specialChars, string(char)) {
			return SpecialRuleError
		}
	}

	return nil
}
