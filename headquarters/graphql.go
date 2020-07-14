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
		crew: Crew!
	}

	type Delivery {
		uuid: String!
		numberOfPackages: Int!
		deliveryDate: String!
	}

	type Crew {
		crewName: String!
		crewMembers: [CrewMember!]!
	}

	type CrewMember {
		name: String!
		role: Role!
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

	enum Role {
		PILOT
		GUNMAN
		NAVIGATOR
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
