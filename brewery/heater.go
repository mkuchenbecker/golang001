package brewery

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	model "github.com/golang001/brewery/model/gomodel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// HeaterServer implements switch.
type HeaterServer struct {
}

func (s *HeaterServer) On(ctx context.Context, req *model.OnRequest) (*model.OnResponse, error) {
	fmt.Printf("On: %s - %s\n", "req.ID", time.Now().String())
	return &model.OnResponse{}, nil
}

func (s *HeaterServer) Off(ctx context.Context, req *model.OffRequest) (*model.OffResponse, error) {
	fmt.Printf("Off: %s - %s\n", "req.ID", time.Now().String())
	return &model.OffResponse{}, nil
}

func main() {
	port := ":8090"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serve := grpc.NewServer()
	model.RegisterSwitchServer(serve, &HeaterServer{})
	// Register reflection service on gRPC server.
	reflection.Register(serve)
	if err := serve.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
