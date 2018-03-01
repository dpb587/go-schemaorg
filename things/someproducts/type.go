package someproducts

import "github.com/dpb587/go-schemaorg"

// // A placeholder for multiple similar products of the same kind.
var Type = schemaorg.NewDataType("http://schema.org", "SomeProducts")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
