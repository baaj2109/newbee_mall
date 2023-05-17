package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // env
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // address
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // db type:mysql(default)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // oss type
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // use multipoint
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimitCount" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`
}
