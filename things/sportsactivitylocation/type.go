package sportsactivitylocation

import "github.com/dpb587/go-schemaorg"

// // A sports location, such as a playing field.
var Type = schemaorg.NewDataType("http://schema.org", "SportsActivityLocation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
