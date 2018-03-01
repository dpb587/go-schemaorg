package state

import "github.com/dpb587/go-schemaorg"

// // A state or province of a country.
var Type = schemaorg.NewDataType("http://schema.org", "State")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
