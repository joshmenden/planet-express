package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/joshmenden/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func getShip(client pb.PlanetExpressClient) (pb.Ship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetShip(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return pb.Ship{}, err
	}

	return *resp.Ship, nil
}

func getCrew(client pb.PlanetExpressClient) (pb.Crew, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetCrew(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.GetCrew(_) = _, %v: ", client, err)
		return pb.Crew{}, err
	}

	return *resp.Crew, nil
}

func listDeliveries(client pb.PlanetExpressClient) ([]*pb.Delivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.ListDeliveries(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.ListDeliveries(_) = _, %v: ", client, err)
		return nil, err
	}

	return resp.Deliveries, nil
}

func getDelivery(client pb.PlanetExpressClient) (pb.Delivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &pb.GetDeliveryRequest{
		Uuid: "1fd2567c-78fa-457b-a9dd-176826087058",
	}

	resp, err := client.GetDelivery(ctx, request)

	if err != nil {
		log.Fatalf("%v.GetDelivery(_) = _, %v: ", client, err)
		return pb.Delivery{}, err
	}

	return *resp.Delivery, nil
}

func main() {
	log.Println("Planet Express Headquarters")

	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)
	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)

	ship, _ := getShip(client)
	// crew, _ := getCrew(client)
	// deliveries, _ := listDeliveries(client)
	// delivery, _ := getDelivery(client)
	testJSON, err := (&jsonpb.Marshaler{OrigName: true}).MarshalToString(&ship)

	log.Println(testJSON)
	// log.Println(ship)
	// log.Println(crew)
	// log.Println(deliveries)
	// log.Println(delivery)
}
