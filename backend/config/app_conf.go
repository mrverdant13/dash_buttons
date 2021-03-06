package config

type appConf struct {
	GraphQLServerConf GraphQLServerConf `mapstructure:"gql_server_conf"`
	DbConf            DbConf            `mapstructure:"db_conf"`
	AdminUser         AdminUser         `mapstructure:"admin_user"`
	JWTConf           JWTConf           `mapstructure:"jwt_conf"`
}
