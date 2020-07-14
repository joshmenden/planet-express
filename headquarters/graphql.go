package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	pb "github.com/joshmenden/planet-express/ship/pkg/planetexpress"
)

const schemaString = `
	schema {
		query: Query
	}
	type Ship {
		name: String!
		location: String!
		fuelLevel: FuelLevel!
		deliveries: [Delivery!]!
	}

	type Delivery {
		uuid: String!
		numberOfPackages: Int!
		deliveryDate: String!
	}

	type Query {
		ship: Ship!
		deliveries: [Delivery!]!
		delivery(uuid: String!): Delivery!
	}

	enum FuelLevel {
		EMPTY
		LOW
		FULL
	}
`

type RootResolver struct{}

func (r *RootResolver) Ship() (pb.Ship, error) {
	client, conn := getPbClient()
	defer conn.Close()
	pbShip, _ := getShip(client)
	return pbShip, nil
}

func (r *RootResolver) Deliveries() ([]*pb.Delivery, error) {
	client, conn := getPbClient()
	defer conn.Close()
	deliveries, _ := listDeliveries(client)
	return deliveries, nil
}

func (r *RootResolver) Delivery(args struct{ UUID string }) (pb.Delivery, error) {
	client, conn := getPbClient()
	defer conn.Close()
	delivery, _ := getDelivery(client, args.UUID)
	return delivery, nil
}

var (
	opts   = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	Schema = graphql.MustParseSchema(schemaString, &RootResolver{}, opts...)
)
