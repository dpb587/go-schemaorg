package casino

import "github.com/dpb587/go-schemaorg"

// // A casino.
var Type = schemaorg.NewDataType("http://schema.org", "Casino")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
