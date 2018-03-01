package movierentalstore

import "github.com/dpb587/go-schemaorg"

// // A movie rental store.
var Type = schemaorg.NewDataType("http://schema.org", "MovieRentalStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
