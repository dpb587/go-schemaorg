package apireference

import "github.com/dpb587/go-schemaorg"

// // Reference documentation for application programming interfaces (APIs).
var Type = schemaorg.NewDataType("http://schema.org", "APIReference")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
