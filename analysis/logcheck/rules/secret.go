package rules

import (
	"errors"
	"strings"
)

var SecretRuleError = errors.New("log message should not contain any secret data")

type SecretRule struct {
	secretData []string
}

type SecretProvider interface {
	Provide() []string
}

func NewSecretRule(secretProvider SecretProvider) *SecretRule {
	return &SecretRule{
		secretData: secretProvider.Provide(),
	}
}

func (rule *SecretRule) Check(msg string) error {
	msg = strings.ToLower(msg)

	for _, secret := range rule.secretData {
		if strings.Contains(msg, secret) {
			return SecretRuleError
		}
	}

	return nil
}
