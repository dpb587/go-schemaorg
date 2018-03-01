package productmodel

import "github.com/dpb587/go-schemaorg"

// // A datasheet or vendor specification of a product (in the sense of a
// prototypical description).
var Type = schemaorg.NewDataType("http://schema.org", "ProductModel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
