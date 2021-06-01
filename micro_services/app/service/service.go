package service

import (
	"context"
	"fmt"
	"log"
	"micro-services/app/registry"
	"net/http"
)

func Start(ctx context.Context, host, port string, reg registry.Registration,
	registerHandlerFunc func()) (context.Context, error) {
	registerHandlerFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)

	err := registry.RegisterService(reg)

	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started service, press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		fmt.Println("service removed:::::::::::::", fmt.Sprintf("http://%v:%v", host, port))
		err := registry.ShutdownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
