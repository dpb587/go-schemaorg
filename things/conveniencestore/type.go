package conveniencestore

import "github.com/dpb587/go-schemaorg"

// // A convenience store.
var Type = schemaorg.NewDataType("http://schema.org", "ConvenienceStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
