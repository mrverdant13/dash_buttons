package config

// JWTConf holds config data about JWT.
type JWTConf struct {
	SecretKey string `mapstructure:"sectret_key"`
}
