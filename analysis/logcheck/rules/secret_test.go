package rules

import (
	"testing"
)

type MockSecretProvider struct {
	secrets []string
}

func (m *MockSecretProvider) Provide() []string {
	return m.secrets
}

func TestSecretRule_Check(t *testing.T) {
	provider := &MockSecretProvider{
		secrets: []string{"password", "token", "secret", "api_key"},
	}
	rule := NewSecretRule(provider)

	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "no secrets in message",
			msg:     "user logged in successfully",
			wantErr: false,
		},
		{
			name:    "contains password secret",
			msg:     "user password: 12345",
			wantErr: true,
		},
		{
			name:    "contains token secret",
			msg:     "auth token: abc123",
			wantErr: true,
		},
		{
			name:    "contains secret word",
			msg:     "this is a secret message",
			wantErr: true,
		},
		{
			name:    "contains api_key secret upper_case",
			msg:     "API_KEY=12345",
			wantErr: true,
		},
		{
			name:    "case insensitive - mixed case",
			msg:     "user PaSsWoRd: 12345",
			wantErr: true,
		},
		{
			name:    "secret as part of word",
			msg:     "this is a password123 test",
			wantErr: true,
		},
		{
			name:    "empty message",
			msg:     "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecretRule.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
