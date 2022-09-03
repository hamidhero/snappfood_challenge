package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
)

func Init()  {
	if err := gotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}