package main

import (
	"cc-go-micro/ccserver/handler"
	"cc-go-micro/ccserver/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	ccserver "cc-go-micro/ccserver/proto/ccserver"
)

func main() {

	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"192.168.1.105:8500",}
	})

	// New Service
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.ccserver"),
		micro.Address(":12345"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	ccserver.RegisterCcserverHandler(service.Server(), new(handler.Ccserver))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.ccserver", service.Server(), new(subscriber.Ccserver))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
