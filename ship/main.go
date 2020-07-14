package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/joshmenden/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const port = 10000

var crew = pb.GetCrewResponse{
	Crew: &pb.Crew{
		NumberOfCrewmates: 5,
		CrewName:          "flying dutchman's squadron",
	},
}

type planetExpressShipServer struct {
	pb.UnimplementedPlanetExpressServer
}

func newServer() *planetExpressShipServer {
	s := &planetExpressShipServer{}
	return s
}

func (s *planetExpressShipServer) GetShip(ctx context.Context, empty *empty.Empty) (*pb.GetShipResponse, error) {
	return &pb.GetShipResponse{
		Ship: &pb.Ship{
			Name:      "the millenium falcon",
			Location:  "tatooine",
			FuelLevel: pb.Ship_FULL,
		},
	}, nil
}

func (s *planetExpressShipServer) GetCrew(ctx context.Context, empty *empty.Empty) (*pb.GetCrewResponse, error) {
	return &crew, nil
}

func (s *planetExpressShipServer) GetDelivery(ctx context.Context, request *pb.GetDeliveryRequest) (*pb.GetDeliveryResponse, error) {
	return &pb.GetDeliveryResponse{
		Delivery: &pb.Delivery{
			Uuid:             request.GetUuid(),
			NumberOfPackages: 50,
			DeliveryDate:     "2020-09-15 00:00:00 -0600",
		},
	}, nil
}

func (s *planetExpressShipServer) ListDeliveries(ctx context.Context, empty *empty.Empty) (*pb.ListDeliveriesResponse, error) {
	return &pb.ListDeliveriesResponse{
		Deliveries: []*pb.Delivery{
			&pb.Delivery{
				NumberOfPackages: 3,
				DeliveryDate:     "2020-07-15 00:00:00 -0600",
			},
			&pb.Delivery{
				NumberOfPackages: 1,
				DeliveryDate:     "2020-07-17 00:00:00 -0600",
			},
			&pb.Delivery{
				NumberOfPackages: 12,
				DeliveryDate:     "2020-08-17 00:00:00 -0600",
			},
		},
	}, nil
}

func main() {
	log.Println("Planet Express Ship")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPlanetExpressServer(grpcServer, newServer())

	log.Printf("Ship listening on localhost:%d\n", port)
	grpcServer.Serve(lis)
}
