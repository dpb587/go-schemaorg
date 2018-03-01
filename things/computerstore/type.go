package computerstore

import "github.com/dpb587/go-schemaorg"

// // A computer store.
var Type = schemaorg.NewDataType("http://schema.org", "ComputerStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
