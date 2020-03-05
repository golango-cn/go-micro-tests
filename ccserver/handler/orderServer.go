package handler

import (
	"cc-go-micro/ccserver/proto/orderserver"
	"context"
	"github.com/micro/go-micro/util/log"
	"math/rand"
)

type OrderServer struct {
}

// CreateOrder
func (o OrderServer) CreateOrder(ctx context.Context, requ *orderserver.CreateOrderRequ, resp *orderserver.CreateOrderResp) error {

	orderId := rand.Int63n(10000)

	log.Info("创建订单：", requ.UserId, requ.UserId)
	log.Info("创建成功：", orderId)

	resp.OrderNumber = orderId
	resp.OrderAmount = 100

	return nil
}

// GetOrder
func (o OrderServer) GetOrder(ctx context.Context, requ *orderserver.GetOrderRequ, resp *orderserver.GetOrderResp) error {

	resp.OrderNumber = requ.OrderNumber
	resp.OrderTitle = "Test title"

	return nil
}
