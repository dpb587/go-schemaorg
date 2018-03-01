package service

import "github.com/dpb587/go-schemaorg"

// // A service provided by an organization, e.g. delivery service, print services,
// etc.
var Type = schemaorg.NewDataType("http://schema.org", "Service")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
