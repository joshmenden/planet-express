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
		# Generic greeting, e.g. "Hello, world!":
		greet: String!
		# Customized greeting, e.g. "Hello, Johan!":
		greetPerson(person: String!): String!
		# More customized greeting, e.g. "Good morning, Johan!":
		greetPersonTimeOfDay(person: String!, timeOfDay: TimeOfDay!): String!
	}
	# Enumerate times of day:
	enum TimeOfDay {
		MORNING
		AFTERNOON
		EVENING
	}
`

type RootResolver struct{}

func (*RootResolver) Greet() string {
	return "Hello, world!"
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
