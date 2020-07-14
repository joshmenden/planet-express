package main

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

const schemaString = `
	schema {
		query: Query
	}

	type Query {
		greet: String!
		greetPerson(person: String!): String!
		greetPersonTimeOfDay(person: String!, timeOfDay: TimeOfDay!): String!
	}

	enum TimeOfDay {
		MORNING
		AFTERNOON
		EVENING
	}
`

type RootResolver struct{}

func (*RootResolver) Greet() string {
	return "Hello, world 2.0!"
}

func (*RootResolver) GreetPerson(args struct{ Person string }) string {
	return fmt.Sprintf("Hello, %s!", args.Person)
}

type PersonTimeOfDayArgs struct {
	Person    string // Note that fields need to be exported.
	TimeOfDay string
}

var TimesOfDay = map[string]string{
	"MORNING":   "Good morning",
	"AFTERNOON": "Good afternoon",
	"EVENING":   "Good evening",
}

func (*RootResolver) GreetPersonTimeOfDay(ctx context.Context, args PersonTimeOfDayArgs) string {
	timeOfDay, ok := TimesOfDay[args.TimeOfDay]
	if !ok {
		timeOfDay = "Go to bed"
	}
	return fmt.Sprintf("%s, %s!", timeOfDay, args.Person)
}

var Schema = graphql.MustParseSchema(schemaString, &RootResolver{})
