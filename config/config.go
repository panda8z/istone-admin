package config

type Server struct {
	Jwt    Jwt    `mapsturcture:"jwt" json:"jwt" yaml:"jwt"`
	Mysql  Mysql  `mapsturcture:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap    `mapsturcture:"zap" json:"zap" yaml:"zap"`
	System System `mapsturcture:"system" json:"system" yaml:"system"`
}
