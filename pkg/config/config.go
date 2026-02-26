package config

type Config struct {
	EnabledRules      []string `yaml:"enabled-rules"`
	SensitivePatterns []string `yaml:"sensitive-patterns"`
}

func Default() *Config {
	return &Config{
		EnabledRules:      []string{"lowercase", "english", "no-special", "no-sensitive"},
		SensitivePatterns: []string{"password", "api_key", "token", "secret", "credit_card", "ssn"},
	}
}
