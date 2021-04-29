package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	a := App{}

	fmt.Println("configuration.database >>>", viper.GetString("database"))
	a.Initialize(viper.GetString("database"), viper.GetString("user"), viper.GetString("password"))

	a.Run(viper.GetString("host"))

}
