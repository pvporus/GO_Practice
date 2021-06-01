package main

import (
	"context"
	"fmt"
	stlog "log"
	"micro-services/app/grades"
	"micro-services/app/registry"
	"micro-services/app/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceUrl = serviceAddress

	ctx, err := service.Start(context.Background(), host, port, r, grades.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)

	}
	<-ctx.Done()
	fmt.Println("shutting down the grading service")

}
