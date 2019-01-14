package brewery

import (
	"context"
	"fmt"
	"log"
	"net"

	model "github.com/golang001/brewery/model/gomodel"
	"github.com/golang001/brewery/rpi/gpio"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// HeaterServer implements switch.
type ThermometerServer struct {
	ctrl    gpio.Controller
	address gpio.TemperatureAddress
}

func (s *ThermometerServer) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	temp, err := s.ctrl.ReadTemperature(s.address)
	return &model.GetResponse{Temperature: temp}, err
}

func StartThermometer(port int, address string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serve := grpc.NewServer()
	addr, err := gpio.NewTemperatureAddress(address, &gpio.DefaultSensorArray{})
	if err != nil {
		log.Fatalf("failed to read address: %v", err)
	}
	model.RegisterThermometerServer(serve, &ThermometerServer{address: addr})
	// Register reflection service on gRPC server.
	reflection.Register(serve)
	if err := serve.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
