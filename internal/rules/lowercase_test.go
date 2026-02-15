package rules

import "testing"

func TestLowercaseRule_Check(t *testing.T) {
	rule := NewLowercaseRule()

	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "starts with lowercase letter",
			msg:     "hello world",
			wantErr: false,
		},
		{
			name:    "starts with uppercase letter",
			msg:     "Hello world",
			wantErr: true,
		},
		{
			name:    "empty string",
			msg:     "",
			wantErr: false,
		},
		{
			name:    "starts with number",
			msg:     "123 hello",
			wantErr: false,
		},
		{
			name:    "starts with special character",
			msg:     "!hello",
			wantErr: false,
		},
		{
			name:    "russian uppercase",
			msg:     "Привет мир",
			wantErr: true,
		},
		{
			name:    "russian lowercase",
			msg:     "привет мир",
			wantErr: false,
		},
		{
			name:    "single uppercase letter",
			msg:     "A",
			wantErr: true,
		},
		{
			name:    "single lowercase letter",
			msg:     "a",
			wantErr: false,
		},
		{
			name:    "starts with space",
			msg:     " hello",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("LowercaseRule.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
