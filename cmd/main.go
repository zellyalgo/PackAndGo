package main

import (
	"fmt"

	"PackAndGo.com/trip/V1/pkg/server"
	"github.com/spf13/viper"
)

func main() {
    viper.AddConfigPath("resources")
    viper.SetConfigName("config")
    viper.SetConfigType("env")
    viper.AutomaticEnv()
    err := viper.ReadInConfig()
    if err != nil {
    	fmt.Printf("%v\n", err)
    }

	app := server.NewServer()
	app.Start()
}