package main

import (
	"cc-go-micro/ccclient/handler"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	"net/http"
)

func main() {

	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// create new web service
	service := web.NewService(
		web.Registry(reg),
		web.Name("go.micro.web.ccclient"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/ccclient/call", handler.CcclientCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
