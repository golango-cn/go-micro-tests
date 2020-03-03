package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	ccserver "cc-go-micro/ccserver/proto/ccserver"
)

type Ccserver struct{}

func (e *Ccserver) Handle(ctx context.Context, msg *ccserver.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *ccserver.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
