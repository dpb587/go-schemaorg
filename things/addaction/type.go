package addaction

import "github.com/dpb587/go-schemaorg"

// // The act of editing by adding an object to a collection.
var Type = schemaorg.NewDataType("http://schema.org", "AddAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
