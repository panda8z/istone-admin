package config

type Mysql struct {
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DatabaseName string `mapstructure:"database-name" json:"database-name" yaml:"database-name"`
	LogMode      bool   `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
	MaxIdleConns int    `mapstructure:"max-idle-conn" json:"max-idle-conn" yaml:"max-idle-conn"`
	MaxOpenConns int    `mapstructure:"max-open-conn" json:"max-open-conn" yaml:"max-open-conn"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Url          string `mapstructure:"url" json:"url" yaml:"url"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
}

func (m Mysql) DNS() string {
	return m.User + ":" + m.Password + "@tcp(" + m.User + ")/" + m.DatabaseName + "?" + m.Config
}
