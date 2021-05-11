package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//NOTE:cert and key files geenrated through OpenSSL
	serverStructTLSConfigService()

}

//HTTPS using TLS API
func httpsService() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hi Golang, how are you doing?")
	})

	log.Fatal(http.ListenAndServeTLS(":9004", "localhost.crt", "localhost.key", nil))
}

//HTTPS using Server struct
func serverStructService() {
	serverStruct := &http.Server{
		Addr:    ":9004",
		Handler: nil,
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hi Golang, how are you doing?")
	})

	log.Fatal(serverStruct.ListenAndServeTLS("localhost.crt", "localhost.key"))

}

//HTTPS using Server Struct and TLSConfig struct
func serverStructTLSConfigService() {

	cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")

	if err != nil {
		log.Fatal(err)
	}

	serverStruct := &http.Server{
		Addr:    ":9004",
		Handler: nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hi Golang, how are you doing?")
	})

	log.Fatal(serverStruct.ListenAndServeTLS("", ""))

}
