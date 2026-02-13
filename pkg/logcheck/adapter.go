package logcheck

type SecretProvider struct {
	data []string
}

func (p SecretProvider) Provide() []string {
	return p.data
}

type SpecialProvider struct {
	data []rune
}

func (p SpecialProvider) Provide() []rune {
	return p.data
}
