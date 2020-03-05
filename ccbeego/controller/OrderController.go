package controller

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	btx "github.com/astaxie/beego/context"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"math/rand"
	"strconv"

	//"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"go-micro-tests/ccbeego/proto/orderserver"
	"time"
)

type OrderController struct {
	Client interface{}
	beego.Controller
}

func pub(name string) {

	msg := &broker.Message{
		Header: map[string]string{
			"name": fmt.Sprintf("%s", name),
		},
		Body: []byte(fmt.Sprintf("%s:%s", name, time.Now().String())),
	}
	if err := broker.Publish("notification.submit", msg); err != nil {
		fmt.Printf("[pub] 发布消息：%s", err)
	} else {
		fmt.Printf("[pub] 发布消息: %s", string(msg.Body))
	}
}

// CreateOrder
func (s *OrderController) CreateOrder(ctx *btx.Context) {

	// Test send message by Rabbitmq
	b := rabbitmq.NewBroker(
		broker.Addrs(`amqp://test01:test01@140.143.0.152:5672`),
	)
	b.Init()
	b.Connect()
	b.Publish("notification.submit", &broker.Message{
		Body: []byte(strconv.Itoa(rand.Intn(1000))),
	})

	// OrderServiceClient
	client := s.Client.(orderserver.OrderServerService)
	rsp, err := client.CreateOrder(context.TODO(), &orderserver.CreateOrderRequ{
		OrderTitle: "ThinkPad x270",
		UserId:     10001,
	})
	if err != nil {
		fmt.Fprint(ctx.ResponseWriter, err.Error())
		return
	}

	ctx.Output.SetStatus(200)
	ctx.Output.JSON(rsp, false, true)

}

// GetOrder
func (s *OrderController) GetOrder(ctx *btx.Context) {

	// OrderServiceClient
	client := s.Client.(orderserver.OrderServerService)

	rsp, err := client.GetOrder(context.TODO(), &orderserver.GetOrderRequ{
		OrderNumber: 100020,
	})
	if err != nil {
		fmt.Fprint(ctx.ResponseWriter, err.Error())
		return
	}

	ctx.Output.SetStatus(200)
	ctx.Output.JSON(rsp, false, true)

}
