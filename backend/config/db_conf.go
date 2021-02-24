package config

// DbConf holds config data about the destination SQL database.
type DbConf struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int64  `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Table    string `mapstructure:"table"`
}
