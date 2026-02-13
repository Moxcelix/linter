package rules

import (
	"testing"
)

type MockSpecialSymbolsProvider struct {
	symbols []rune
}

func (m *MockSpecialSymbolsProvider) Provide() []rune {
	return m.symbols
}

func TestSpecialRule_Check(t *testing.T) {
	provider := &MockSpecialSymbolsProvider{
		symbols: []rune{'!', '@', '#', '$', '%', '^', '&', '*'},
	}
	rule := NewSpecialRule(provider)

	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "no special characters",
			msg:     "hello world",
			wantErr: false,
		},
		{
			name:    "contains exclamation mark",
			msg:     "hello! world",
			wantErr: true,
		},
		{
			name:    "contains @ symbol",
			msg:     "user@example.com",
			wantErr: true,
		},
		{
			name:    "empty message",
			msg:     "",
			wantErr: false,
		},
		{
			name:    "contains multiple special characters",
			msg:     "hello!@#$ world",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpecialRule.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
