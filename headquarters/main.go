package main

import (
	"flag"
	"log"

	pb "github.com/joshmenden/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

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
	ship, _ := getShip(client)
	json, err := (&jsonpb.Marshaler{OrigName: true}).MarshalToString(&ship)

	log.Printf(json)

	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)
	// http.Handle("/query", &relay.Handler{Schema: Schema})
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
