package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func main() {
	log.Println("Planet Express Headquarters")

	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)
	http.Handle("/query", DisableCors(&relay.Handler{Schema: Schema}))
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Printf("Listening for GQL Queries on addr: localhost:8080")
}
