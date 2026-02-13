package rules

import "testing"

func TestEnglishRule_Check(t *testing.T) {
	rule := NewEnglishRule()

	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "english only",
			msg:     "hello world",
			wantErr: false,
		},
		{
			name:    "russian characters",
			msg:     "hello мир",
			wantErr: true,
		},
		{
			name:    "empty string",
			msg:     "",
			wantErr: false,
		},
		{
			name:    "mixed text",
			msg:     "Hello, мир!",
			wantErr: true,
		},
		{
			name:    "numbers and symbols",
			msg:     "123!@#",
			wantErr: false,
		},
		{
			name:    "emoji",
			msg:     "hello 🚀",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnglishRule.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
