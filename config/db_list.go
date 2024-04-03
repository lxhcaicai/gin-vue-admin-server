package config

type DsnProvider interface {
	Dsn() string
}

type GeneralDB struct {
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`

	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"` // 数据库名
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // 数据库密码
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   //是否开启全局禁用复数，true表示开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        //数据库引擎，默认InnoDB
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数

}

type SpecializedDB struct {
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
	Disable   bool `mapstructure:"disable" json:"disable" yaml:"disable"`
}
