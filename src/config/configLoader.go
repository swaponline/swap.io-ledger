package config

import (
	"github.com/spf13/viper"
	"log"
)

var PORT string
var SECRET_TOKEN string
var POSTGRESS_URL string
var AGENTS []struct{
	Network  string `mapstructure:"network"`
    BaseUrl  string `mapstructure:"baseUrl"`
    ApiKey   string `mapstructure:"apiKey"`
}

func InitializeConfig() {
	cfgReader := viper.New()

	cfgReader.AddConfigPath(".")
	cfgReader.SetConfigFile(".env.yaml")
	err := cfgReader.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

    if err = cfgReader.UnmarshalKey("port", &PORT); err != nil {
        log.Panicln(err)
    }

    if err = cfgReader.UnmarshalKey("secret_token", &SECRET_TOKEN); err != nil {
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

	if len(SECRET_TOKEN) == 0 {
		log.Panicln("ERROR set secret_token in env")
	}
}
