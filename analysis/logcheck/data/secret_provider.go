package data

import (
	"main/analysis/logcheck/rules"
)

type SecretProvider struct{}

func NewSecretProvider() rules.SecretProvider {
	return &SecretProvider{}
}

func (s *SecretProvider) Provide() []string {
	return []string{"apikey", "api-key", "api_key", "token", "jwt_token", "jwt-token", "secret", "key"}
}
