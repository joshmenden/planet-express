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

var deliveries = pb.ListDeliveriesResponse{
	Deliveries: []*pb.Delivery{
		&pb.Delivery{
			Uuid:             "c7b3e7e3-8049-4741-a1cd-32a70f2165fb",
			NumberOfPackages: 3,
			DeliveryDate:     "2020-07-15 00:00:00 -0600",
		},
		&pb.Delivery{
			Uuid:             "d220d52b-ab27-42df-96ca-1a3693ad3477",
			NumberOfPackages: 1,
			DeliveryDate:     "2020-07-17 00:00:00 -0600",
		},
		&pb.Delivery{
			Uuid:             "1fd2567c-78fa-457b-a9dd-176826087058",
			NumberOfPackages: 12,
			DeliveryDate:     "2020-08-17 00:00:00 -0600",
		},
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
	// search for delivery by given uuid
	var desiredPackage *pb.Delivery
	for _, v := range deliveries.Deliveries {
		if v.GetUuid() == request.GetUuid() {
			desiredPackage = v
		}
	}
	return &pb.GetDeliveryResponse{
		Delivery: desiredPackage,
	}, nil
}

func (s *planetExpressShipServer) ListDeliveries(ctx context.Context, empty *empty.Empty) (*pb.ListDeliveriesResponse, error) {
	return &deliveries, nil
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
