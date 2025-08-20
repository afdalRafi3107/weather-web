package utils

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found")
	}
}