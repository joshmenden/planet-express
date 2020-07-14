package main

import (
	"flag"
	"log"

	pb "github.com/joshmenden/planet-express/ship/pkg/planetexpress"
	"google.golang.org/grpc"
)

func getPbClient() (pb.PlanetExpressClient, *grpc.ClientConn) {
	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := pb.NewPlanetExpressClient(conn)
	return client, conn
}
