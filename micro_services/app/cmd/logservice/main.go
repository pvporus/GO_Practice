package main

import (
	"context"
	"fmt"
	stlog "log"
	"micro-services/app/log"
	"micro-services/app/registry"
	"micro-services/app/service"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"
	serviceAddes := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceUrl = serviceAddes

	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down the log service")

}
