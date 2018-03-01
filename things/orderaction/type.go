package orderaction

import "github.com/dpb587/go-schemaorg"

// // An agent orders an object/product/service to be delivered/sent.
var Type = schemaorg.NewDataType("http://schema.org", "OrderAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
