package main

import (
	graphql "github.com/graph-gophers/graphql-go"
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
		// number_of_packages: Int
		deliveryDate: String!
	}

	type Query {
		ships: [Ship!]!
	}

	enum FuelLevel {
		EMPTY
		LOW
		FULL
	}
`

type Ship struct {
	Name       string
	Location   string
	FuelLevel  string
	Deliveries []Delivery
}

type Delivery struct {
	Uuid             string
	NumberOfPackages int
	DeliveryDate     string
}

var ships = []Ship{
	{
		Name:      "test ship here",
		Location:  "Lehi UT",
		FuelLevel: "FULL",
		Deliveries: []Delivery{
			{
				Uuid: "abc987654321",
				// NumberOfPackages: 4,
				DeliveryDate: "march 3 1990",
			},
		},
	},
}

type RootResolver struct{}

func (r *RootResolver) Ships() ([]Ship, error) {
	return ships, nil
}

var (
	// We can pass an option to the schema so we don’t need to
	// write a method to access each type’s field:
	opts   = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	Schema = graphql.MustParseSchema(schemaString, &RootResolver{}, opts...)
)

// var Schema = graphql.MustParseSchema(schemaString, &RootResolver{})
