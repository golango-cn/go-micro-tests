package main

import (
	"cc-go-micro/ccserver/handler"
	"cc-go-micro/ccserver/proto/orderserver"
	"cc-go-micro/ccserver/proto/userserver"
	"cc-go-micro/ccserver/subscriber"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"

	ccserver "cc-go-micro/ccserver/proto/ccserver"
)

func main() {

	b := rabbitmq.NewBroker(
		broker.Addrs(`amqp://test01:test01@140.143.0.152:5672`),
	)

	b.Init()
	b.Connect()


	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"192.168.1.105:8500",}
	})

	// New Service
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.ccserver"),
		micro.Address(":12345"),
		micro.Version("latest"),
		//micro.Broker(b),
	)

	// Initialise service
	service.Init()

	// Register Handler
	ccserver.RegisterCcserverHandler(service.Server(), new(handler.Ccserver))
	orderserver.RegisterOrderServerHandler(service.Server(), new(handler.OrderServer))
	userserver.RegisterUserServerHandler(service.Server(), new(handler.UserServer))

	b.Subscribe("notification.submit", subscriber.ProtoHandler, broker.Queue("hello2"))

	//Register Subscriber
	micro.RegisterSubscriber("notification.submit", service.Server(), new(subscriber.Ccserver))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
