package eataction

import "github.com/dpb587/go-schemaorg"

// // The act of swallowing solid objects.
var Type = schemaorg.NewDataType("http://schema.org", "EatAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
