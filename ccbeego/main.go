package main

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
	"go-micro-tests/ccbeego/controller"
	"go-micro-tests/ccbeego/proto/ccserver"
	"go-micro-tests/ccbeego/proto/orderserver"
	"log"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"github.com/astaxie/beego"
)

func main() {

	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.1.105:8500",
		}
	})
	service := web.NewService(
		web.Registry(reg),
		web.Address(":8080"),
		web.Name("go.micro.client.web"))
	service.Init()

	// call the backend service
	sayClient := ccserver.NewCcserverService("go.micro.srv.ccserver", client.DefaultClient)

	// Create RESTful handler
	say := &controller.SayController{Client: sayClient,}
	beego.Get("/call/:name", say.Call)
	beego.Get("/stream/:name", say.Stream)


	orderClient := orderserver.NewOrderServerService("go.micro.srv.ccserver", client.DefaultClient)
	order := &controller.OrderController{Client: orderClient, }

	beego.Get("/order/create", order.CreateOrder)
	beego.Get("/order/:id", order.GetOrder)

	// Register Handler
	service.Handle("/", beego.BeeApp.Handlers)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
