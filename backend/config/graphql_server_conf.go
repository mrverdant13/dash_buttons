package config

import "strconv"

// GraphQLServerConf holds GraphQL server config data.
type GraphQLServerConf struct {
	Port uint64 `mapstructure:"port"`
}

// PortString returns the GraphQL port number as string.
func (r *GraphQLServerConf) PortString() string {
	return strconv.FormatUint(r.Port, 10)
}
