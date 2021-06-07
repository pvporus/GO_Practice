package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	defer logger.Sync()
	//Initialize Viper config
	initializeConfig()

	//Zap looger initialization
	logger = getLogger()

	a := App{}

	logger.Info(fmt.Sprintf("Reading from config : database %s, user %s, password %s", viper.GetString("database"), viper.GetString("user"), viper.GetString("password")))
	a.Initialize(viper.GetString("database"), viper.GetString("user"), viper.GetString("password"))
	logger.Info("Initialization completed...")
	a.Run(viper.GetString("host"))

}

func getLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{viper.GetString("logfile")}
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return logger
}

func initializeConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
