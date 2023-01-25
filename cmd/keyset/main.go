package main

import (
	"os"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("client")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.Set("RIOT_API_KEY", os.Args[1])
	viper.WriteConfig()
}
