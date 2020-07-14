package main

import (
	"context"
	"log"
	"time"

	pb "github.com/joshmenden/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/ptypes/empty"
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

func getDelivery(client pb.PlanetExpressClient, UUID string) (pb.Delivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &pb.GetDeliveryRequest{
		Uuid: UUID,
	}

	resp, err := client.GetDelivery(ctx, request)

	if err != nil {
		log.Fatalf("%v.GetDelivery(_) = _, %v: ", client, err)
		return pb.Delivery{}, err
	}

	return *resp.Delivery, nil
}
