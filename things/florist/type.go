package florist

import "github.com/dpb587/go-schemaorg"

// // A florist.
var Type = schemaorg.NewDataType("http://schema.org", "Florist")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
