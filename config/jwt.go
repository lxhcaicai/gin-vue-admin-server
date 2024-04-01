package config

type JWT struct {
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
}
