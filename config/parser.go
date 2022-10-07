package config

import (
	"github.com/spf13/viper"
	"log"
)

var GITLAB_URL string
var API_KEY string
const GITLAB_METHOD = "GET"

func SetupConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("service.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Config Not Found")
	}
	GITLAB_URL = "https://" + viper.GetString(`host.server`) + "/api/v4/projects/"
	API_KEY = viper.GetString(`host.apikey`)
}