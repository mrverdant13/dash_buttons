package config

import (
	"log"

	"github.com/golobby/container"
	"github.com/spf13/viper"
)

// Init loads and injects all config values from "config.yaml" file located at "path".
//
// Injects:
//
// - GraphQLServerConf
func Init(path string) {
	config := load(path)

	container.Singleton(func() GraphQLServerConf {
		return config.GraphQLServerConf
	})

	container.Singleton(func() DbConf {
		return config.DbConf
	})
}

func load(path string) (config appConf) {

	viper.AddConfigPath(path)

	configFileName := "config"
	viper.SetConfigName(configFileName)

	configType := "yaml"
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return config
}
