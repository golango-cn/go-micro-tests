package controller

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	btx "github.com/astaxie/beego/context"
	"go-micro-tests/ccbeego/proto/ccserver"
)

type SayController struct {
	Client interface{}
	beego.Controller
}

// Call
func (s *SayController) Call(ctx *btx.Context) {

	name := ctx.Input.Param(":name")
	ccclientClient := s.Client.(ccserver.CcserverService)

	rsp, err := ccclientClient.Call(context.TODO(), &ccserver.Request{
		Name: name,
	})
	if err != nil {
		fmt.Fprint(ctx.ResponseWriter, err.Error())
		return
	}

	ctx.Output.SetStatus(200)
	ctx.Output.JSON(rsp, false, true)

}

// Stream
func (s *SayController) Stream(ctx *btx.Context) {

	ccclientClient := s.Client.(ccserver.CcserverService)

	rsp, err := ccclientClient.Stream(context.TODO(), &ccserver.StreamingRequest{
		Count: 100,
	})
	if err != nil {
		fmt.Fprint(ctx.ResponseWriter, err.Error())
		return
	}

	ctx.Output.SetStatus(200)
	ctx.Output.JSON(rsp, false, true)
}
