package country

import "github.com/dpb587/go-schemaorg"

// // A country.
var Type = schemaorg.NewDataType("http://schema.org", "Country")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
