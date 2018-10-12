package main

import (
	"log"

	protos "github.com/golang001/deltav/protos"
	"github.com/golang001/deltav/shuttle"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewWorldModelClient(conn)

	vessel := protos.Vessel{
		Sensors: &protos.SensorSystem{
			Sensors: []*protos.Sensor{
				&protos.Sensor{
					SensorType:    protos.Sensor_PASSIVE,
					RadiationType: protos.RadiationType_GAMMA,
				},
			},
		},
	}
	playerShuttle := shuttle.NewPlayerShuttle(vessel, c)
	r, err := c.Register(ctx, &protos.RegisterRequest{})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("Greeting: %s", r.Effect)
}
