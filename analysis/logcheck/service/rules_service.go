package service

import "main/analysis/logcheck/rules"

type Callback func(err error)

type RulesService struct {
	rules rules.Rules
}

func NewRulesService(rules rules.Rules) *RulesService {
	return &RulesService{
		rules: rules,
	}
}

func (s *RulesService) CheckRules(msg string, callback Callback) {
	for _, rule := range s.rules {
		if err := rule.Check(msg); err != nil {
			callback(err)
		}
	}
}
