package config

type Jwt struct {
	BufferTime  int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
	SignKey     string `mapstructure:"sign-key" json:"sign-key" yaml:"sign-key"`
}
