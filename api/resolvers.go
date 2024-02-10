package api

import "github.com/graphql-go/graphql"

func getPersonByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if ok {
		// Fetch person with the given id from the database
	}
	_ = id
	return nil, nil // Replace with actual data fetching logic
}
