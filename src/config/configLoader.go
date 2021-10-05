package config

import (
	"github.com/spf13/viper"
	"log"
)

var POSTGRESS_URL string
var AGENTS []struct{
	Network  string `mapstructure:"network"`
    BaseUrl  string `mapstructure:"baseUrl"`
    Apikey   string `mapstructure:"apiKey"`
}

func InitializeConfig() {
	cfgReader := viper.New()

	cfgReader.AddConfigPath(".")
	cfgReader.SetConfigFile(".env.yaml")
	err := cfgReader.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

	if err = cfgReader.UnmarshalKey("agents", &AGENTS); err != nil {
		log.Panicln(err)
	}

	if err = cfgReader.UnmarshalKey("postgres.url", &POSTGRESS_URL); err != nil {
		log.Panicln(err)
	}
	if len(POSTGRESS_URL) == 0 {
		log.Panicln("ERROR set postgres url(postgres.url) in env")
	}
}