package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`

	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`

	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
