package sportsclub

import "github.com/dpb587/go-schemaorg"

// // A sports club.
var Type = schemaorg.NewDataType("http://schema.org", "SportsClub")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
