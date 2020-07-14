package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/graph-gophers/graphql-go/relay"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func main() {
	log.Println("Planet Express Headquarters")
	client, conn := getPbClient()
	defer conn.Close()

	ship, _ := getShip(client)
	json, _ := (&jsonpb.Marshaler{OrigName: true}).MarshalToString(&ship)

	log.Printf(json)
	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)
	http.Handle("/query", &relay.Handler{Schema: Schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
