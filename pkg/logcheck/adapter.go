package logcheck

type SecretProvider struct {
	data []string
}

func (p SecretProvider) Provide() []string {
	return p.data
}

type SpecialProvider struct {
	data []string
}

func (p SpecialProvider) Provide() []string {
	return p.data
}
