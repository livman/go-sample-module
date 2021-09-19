package gosamplemodule

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/livman/go-sample-module/pingpong"
)

// Implement pb.PingPongServer
type PingPongServerImpl struct {
}

func (s *PingPongServerImpl) StartPing(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	fmt.Println("Ping Received")

	resp := pb.Pong{
		Id:      ping.Id,
		Message: ping.Message,
	}

	return &resp, nil
}

func GetPb() pb {
	return pb
}
