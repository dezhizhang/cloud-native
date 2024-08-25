package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

func main() {
	var serverConfig ServerConfig
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	err := v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(serverConfig.Name)
}
