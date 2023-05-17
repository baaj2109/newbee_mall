package config

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // host
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // port
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // config
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // database name
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // user name
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // password
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // max idle conns
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // max open conns
	LogMode      string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // log mode
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // log zap
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
