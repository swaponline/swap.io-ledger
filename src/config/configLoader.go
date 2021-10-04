package config

import (
	"github.com/spf13/viper"
	"log"
)

var POSTGRESS_URL string
var AGENTS []struct{
	Network  string `mapstructure:"mode"`
    BaseUrl  string `mapstructure:"mode"`
    Apikey   string `mapstructure:"mode"`
}

func InitializeConfig() {
	viper := viper.New()

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(viper.GetString("postgres.url"))
}