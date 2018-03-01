package steeringpositionvalue

import "github.com/dpb587/go-schemaorg"

// // A value indicating a steering position.
var Type = schemaorg.NewDataType("http://schema.org", "SteeringPositionValue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
