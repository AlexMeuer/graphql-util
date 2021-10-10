package gql

import "github.com/hasura/go-graphql-client"

type (
	// Vars is shorthand for map[string]interface{}, intended to make
	// declaring variables for GraphQL operations more concise.
	Vars    map[string]interface{}
	Boolean graphql.Boolean
	Float   graphql.Float
	ID      graphql.ID
	Int     graphql.Int
	String  graphql.String
	uuid    string
	jsonb   map[string]interface{}
	bigint  int64
)

func UUID(s string) uuid {
	if s == "" {
		return uuid("00000000-0000-0000-0000-000000000000")
	}
	return uuid(s)
}
func JSONB(d map[string]interface{}) jsonb {
	return jsonb(d)
}
func BigInt(n int64) bigint {
	return bigint(n)
}
