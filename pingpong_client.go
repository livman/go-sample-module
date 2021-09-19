package gosamplemodule

import (
	"context"
	"fmt"
	"time"

	pb "github.com/livman/go-sample-module/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendPing(client pb.PingPongClient) (*pb.Pong, error) {
	// Timeout 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ping := pb.Ping{
		Id:      1,
		Message: "Ping",
	}
	pong, err := client.StartPing(ctx, &ping)
	statusCode := status.Code(err)
	if statusCode != codes.OK {
		return nil, err
	}
	fmt.Printf("Pong: %d, statusCode: %s\n", pong.Id, statusCode.String())
	return pong, err
}
