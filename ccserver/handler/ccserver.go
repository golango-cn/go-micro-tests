package handler

import (
	"context"
	"strconv"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	ccserver "cc-go-micro/ccserver/proto/ccserver"
)

type Ccserver struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Ccserver) Call(ctx context.Context, req *ccserver.Request, rsp *ccserver.Response) error {
	log.Info("Received Ccserver.Call request")
	nano := time.Now().UnixNano()

	rsp.Msg = "Hello " + req.Name + "/" + strconv.Itoa(int(nano))
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Ccserver) Stream(ctx context.Context, req *ccserver.StreamingRequest, stream ccserver.Ccserver_StreamStream) error {
	log.Infof("Received Ccserver.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&ccserver.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Ccserver) PingPong(ctx context.Context, stream ccserver.Ccserver_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&ccserver.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
