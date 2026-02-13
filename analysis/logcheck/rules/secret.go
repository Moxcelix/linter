package rules

import (
	"errors"
	"strings"
)

var SecretRuleError = errors.New("log message should not contain any secret data")

var secretData = []string{"apikey", "api-key", "api_key", "token", "jwt_token", "jwt-token", "secret", "key"}

func CheckSecretRule(msg string) error {
	msg = strings.ToLower(msg)

	for _, secret := range secretData {
		if strings.Contains(msg, secret) {
			return SecretRuleError
		}
	}

	return nil
}
