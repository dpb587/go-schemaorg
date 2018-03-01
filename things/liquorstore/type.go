package liquorstore

import "github.com/dpb587/go-schemaorg"

// // A shop that sells alcoholic drinks such as wine, beer, whisky and other
// spirits.
var Type = schemaorg.NewDataType("http://schema.org", "LiquorStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
