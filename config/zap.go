package config

type Zap struct {
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	LinkName      string `mapstructure:"link-name" json:"link-name" yaml:"link-name"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"`
}
