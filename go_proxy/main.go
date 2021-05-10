package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	http.HandleFunc("/", handleRequestAndRedirect)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Handles requests by parsing the custom header (used ModHeader browser plugin with custom header)
//Forward to proxy server
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {

	value, _ := req.Header["Custurl"]

	targeUrl := value[0]

	serveReverseProxy(targeUrl, res, req)

}

//Reverse proxy to forward requests based on the custom header url
//Ex: custUrl: http://gdoc.org request: http://localhost:8080/github.com/stretchr/testify/assert
//Forward the request to the target url: http://godoc.org/github.com/stretchr/testify/assert
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forward-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)

}
