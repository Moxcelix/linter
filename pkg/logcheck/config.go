package logcheck

type Config struct {
	Rules   RulesConfig    `json:"rules" yaml:"rules"`
	Loggers []LoggerConfig `json:"loggers" yaml:"loggers"`
}

type RulesConfig struct {
	English   EnglishRuleConfig   `json:"english" yaml:"english"`
	Lowercase LowercaseRuleConfig `json:"lowercase" yaml:"lowercase"`
	Secret    SecretRuleConfig    `json:"secret" yaml:"secret"`
	Special   SpecialRuleConfig   `json:"special" yaml:"special"`
}

type EnglishRuleConfig struct {
	Enabled bool `json:"enabled" yaml:"enabled"`
}

type LowercaseRuleConfig struct {
	Enabled bool `json:"enabled" yaml:"enabled"`
}

type SecretRuleConfig struct {
	Enabled bool     `json:"enabled" yaml:"enabled"`
	Words   []string `json:"words" yaml:"words"`
}

type SpecialRuleConfig struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Chars   []rune `json:"chars" yaml:"chars"`
}

type LoggerConfig struct {
	PkgName   string            `json:"pkg_name" yaml:"pkg_name"`
	LoggerObj []LoggerObjConfig `json:"logger_obj" yaml:"logger_obj"`
	Funcs     []string          `json:"funcs" yaml:"funcs"`
}

type LoggerObjConfig struct {
	Name    string   `json:"name" yaml:"name"`
	Methods []string `json:"methods" yaml:"methods"`
}

func DefaultConfig() *Config {
	return &Config{
		Rules: RulesConfig{
			English: EnglishRuleConfig{
				Enabled: true,
			},
			Lowercase: LowercaseRuleConfig{
				Enabled: true,
			},
			Secret: SecretRuleConfig{
				Enabled: true,
				Words: []string{
					"password", "passwd",
					"token", "secret", "key",
					"api_key", "apikey",
				},
			},
			Special: SpecialRuleConfig{
				Enabled: true,
				Chars: []rune{
					'!', '@', '#', '$', '%', '^', '&', '*',
				},
			},
		},
		Loggers: []LoggerConfig{
			{
				PkgName: "go.uber.org/zap",
				LoggerObj: []LoggerObjConfig{
					{
						Name: "Logger",
						Methods: []string{
							"Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
							"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
							"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
							"With", "Named", "WithOptions", "Core", "Check", "Sugar",
						},
					},
					{
						Name: "SugaredLogger",
						Methods: []string{
							"Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
							"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
							"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
							"With", "Named", "Desugar",
						},
					},
				},
			},
			{
				PkgName: "log/slog",
				LoggerObj: []LoggerObjConfig{
					{
						Name: "Logger",
						Methods: []string{
							"Debug", "Info", "Warn", "Error",
							"DebugContext", "InfoContext", "WarnContext", "ErrorContext",
							"Log", "LogContext",
							"Enabled", "Handler", "With", "WithGroup",
						},
					},
				},
				Funcs: []string{
					"Debug", "Info", "Warn", "Error",
					"DebugContext", "InfoContext", "WarnContext", "ErrorContext",
					"Log", "LogContext",
					"Default", "SetDefault", "New", "NewJSONHandler", "NewTextHandler",
					"With", "NewRecord", "NewLogLogger",
				},
			},
		},
	}
}
