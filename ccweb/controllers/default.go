package controllers

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-micro-tests/ccweb/proto/ccserver"

	"github.com/micro/go-micro/v2"
)

type MainController struct {
	beego.Controller
}

var ccclientClient ccserver.CcserverService

func init() {

	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.1.105:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.client.web"))
	service.Init()

	// call the backend service
	ccclientClient = ccserver.NewCcserverService("go.micro.srv.ccserver", service.Client())

	fmt.Println("123")
}

func (c *MainController) Get() {

	rsp, err := ccclientClient.Call(context.TODO(), &ccserver.Request{
		Name: "你好",
	})
	if err != nil {
		fmt.Fprint(c.Ctx.ResponseWriter, err.Error())
		return
	}

	c.Data["json"] = rsp
	c.ServeJSON()

}
