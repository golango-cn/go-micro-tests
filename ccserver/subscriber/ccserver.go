package subscriber

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"time"

	ccserver "cc-go-micro/ccserver/proto/ccserver"
)

type Ccserver struct{}

func ProtoHandler(e broker.Event) error {
	fmt.Println("Messages from the queue", e.Topic(), e.Ack(), string(e.Message().Body))
	return nil
}

func (e *Ccserver) Handle(ctx context.Context, msg *ccserver.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	time.Sleep(10 * time.Second)
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *ccserver.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
