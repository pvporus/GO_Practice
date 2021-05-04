package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//NOTE:cert and key files geenrated through OpenSSL
	//Simple TLS Server with http
	httpsService()
	//TLS server with http.Server struct
	/*serverStructService()*/
	//TLS server with http.Server struct and TLSConfig struct
	/*serverStructTLSConfigService()*/

}
func httpsService() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hi Golang, how are you doing?")
	})

	log.Fatal(http.ListenAndServeTLS(":9004", "localhost.crt", "localhost.key", nil))
}

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
