package rules

import "go.uber.org/fx"

type Rule interface {
	Check(msg string) error
}

type Rules []Rule

func NewRules(
	englishRule *EnglishRule,
	lowercaseRule *LowercaseRule,
	secretRule *SecretRule,
	specialRule *SpecialRule,
) Rules {
	return Rules{
		englishRule,
		lowercaseRule,
		secretRule,
		specialRule,
	}
}

var Module = fx.Options(
	fx.Provide(NewEnglishRule),
	fx.Provide(NewLowercaseRule),
	fx.Provide(NewSecretRule),
	fx.Provide(NewSpecialRule),
	fx.Provide(NewRules),
)
