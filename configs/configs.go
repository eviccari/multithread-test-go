package configs

import (
	"log"

	"github.com/spf13/viper"
)

const (
	ConfigPath1 = "./cmd/env"
	ConfigPath2 = "./env"
	ConfigPath3 = "."
	ConfigType  = "json"
	ConfigName  = ".env"
)

var GithubUsersQuantity int
var MultiThreadSize int
var GithubAPIMaxPageSize int

func init() {
	viper.AddConfigPath(ConfigPath1)
	viper.AddConfigPath(ConfigPath2)
	viper.AddConfigPath(ConfigPath3)
	viper.SetConfigType(ConfigType)
	viper.SetConfigName(ConfigName)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error on load job configurations: %s", err.Error())
	}

	GithubUsersQuantity = viper.GetInt("GithubUsersQuantity")
	MultiThreadSize = viper.GetInt("MultiThreadSize")
	GithubAPIMaxPageSize = viper.GetInt("GithubAPIMaxPageSize")
}
