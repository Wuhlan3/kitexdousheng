package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	//workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
