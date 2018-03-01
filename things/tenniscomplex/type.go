package tenniscomplex

import "github.com/dpb587/go-schemaorg"

// // A tennis complex.
var Type = schemaorg.NewDataType("http://schema.org", "TennisComplex")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
