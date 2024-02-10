package api

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"person": &graphql.Field{
			Type: personType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: getPersonByID,
		},
		// Add other queries here
	},
})

// Define your root mutation similarly if needed

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation, // Uncomment if you have mutations
	},
)
