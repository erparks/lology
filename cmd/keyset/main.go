package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	fileName := "client"
	extension := "json"
	fullname := fileName + "." + extension

	if _, err := os.Stat(fullname); err != nil {
		if f, err := os.Create(fullname); err != nil {
			log.Fatal(err)
		} else {
			f.Close()
		}
	}

	viper.SetConfigName("client")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.Set("RIOT_API_KEY", os.Args[1])
	viper.WriteConfig()
}
