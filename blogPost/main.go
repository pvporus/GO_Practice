package main

import (
	"blogPost/api"
	_ "net/http/pprof"
)

func main() {
	api.Run()
}
