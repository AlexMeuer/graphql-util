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

// UUID casts a string to a uuid (The Hasura UUID type is lowercase).
func UUID(s string) uuid {
	if s == "" {
		return uuid("00000000-0000-0000-0000-000000000000")
	}
	return uuid(s)
}

// JSONB casts a map to a jsonb (json binary) object (The Hasura JSONB type is lowercase).
func JSONB(d map[string]interface{}) jsonb {
	return jsonb(d)
}

// BigInt casts an int64 to a bigint (The Hasura BigInt type is lowercase).
func BigInt(n int64) bigint {
	return bigint(n)
}
