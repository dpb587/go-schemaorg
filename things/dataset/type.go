package dataset

import "github.com/dpb587/go-schemaorg"

// // A body of structured information describing some topic(s) of interest.
var Type = schemaorg.NewDataType("http://schema.org", "Dataset")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
