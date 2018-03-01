package cemetery

import "github.com/dpb587/go-schemaorg"

// // A graveyard.
var Type = schemaorg.NewDataType("http://schema.org", "Cemetery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
