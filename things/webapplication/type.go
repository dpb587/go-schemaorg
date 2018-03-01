package webapplication

import "github.com/dpb587/go-schemaorg"

// // Web applications.
var Type = schemaorg.NewDataType("http://schema.org", "WebApplication")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
