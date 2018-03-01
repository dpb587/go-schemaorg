package sportsorganization

import "github.com/dpb587/go-schemaorg"

// // Represents the collection of all sports organizations, including sports
// teams, governing bodies, and sports associations.
var Type = schemaorg.NewDataType("http://schema.org", "SportsOrganization")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
