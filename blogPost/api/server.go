package api

import (
	"blogPost/api/bootstrapping"
	"blogPost/api/controllers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

//Run the http server to listen and serve the requests based on the routs
func Run() {
	//Close the DB once done
	defer server.DB.Close()

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	bootstrapping.Load(server.DB)

	server.Run(os.Getenv("APP_PORT"))

}

//Loading environment variables without Docker
func loadEnv() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}
