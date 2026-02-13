package data

import "main/analysis/logcheck/rules"

type SpecialProvider struct {
}

func NewSpecialProvider() rules.SpecialSymbolsProvider {
	return &SpecialProvider{}
}

func (s *SpecialProvider) Provide() []rune {
	return []rune{'!', '?', '.', ',', ':', ';', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '=', '{', '}', '[', ']', '|', '\\', '/', '<', '>', '~', '`'}
}
