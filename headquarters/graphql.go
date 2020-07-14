package main

import (
	graphql "github.com/graph-gophers/graphql-go"
)

const schemaString = `
	schema {
		query: Query
	}
	type User {
		userID: ID!
		username: String!
		emoji: String!
		notes: [Note!]!
	}

	type Note {
		noteID: ID!
		data: String!
	}

	type Query {
		users: [User!]!
	}
`

type User struct {
	UserID   graphql.ID
	Username string
	Emoji    string
	Notes    []Note
}

type Note struct {
	NoteID graphql.ID
	Data   string
}

var users = []User{
	{
		UserID:   graphql.ID("u-001"),
		Username: "nyxerys",
		Emoji:    "🇵🇹",
		Notes: []Note{
			{NoteID: "n-001", Data: "Olá Mundo!"},
			{NoteID: "n-002", Data: "Olá novamente, mundo!"},
			{NoteID: "n-003", Data: "Olá, escuridão!"},
		},
	}, {
		UserID:   graphql.ID("u-002"),
		Username: "rdnkta",
		Emoji:    "🇺🇦",
		Notes: []Note{
			{NoteID: "n-004", Data: "Привіт Світ!"},
			{NoteID: "n-005", Data: "Привіт ще раз, світ!"},
			{NoteID: "n-006", Data: "Привіт, темрява!"},
		},
	}, {
		UserID:   graphql.ID("u-003"),
		Username: "username_ZAYDEK",
		Emoji:    "🇺🇸",
		Notes: []Note{
			{NoteID: "n-007", Data: "Hello, world!"},
			{NoteID: "n-008", Data: "Hello again, world!"},
			{NoteID: "n-009", Data: "Hello, darkness!"},
		},
	},
}

type RootResolver struct{}

func (r *RootResolver) Users() ([]User, error) {
	return users, nil
}

var (
	// We can pass an option to the schema so we don’t need to
	// write a method to access each type’s field:
	opts   = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	Schema = graphql.MustParseSchema(schemaString, &RootResolver{}, opts...)
)

// var Schema = graphql.MustParseSchema(schemaString, &RootResolver{})
