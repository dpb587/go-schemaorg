package bookstore

import "github.com/dpb587/go-schemaorg"

// // A bookstore.
var Type = schemaorg.NewDataType("http://schema.org", "BookStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
