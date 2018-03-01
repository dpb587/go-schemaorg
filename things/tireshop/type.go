package tireshop

import "github.com/dpb587/go-schemaorg"

// // A tire shop.
var Type = schemaorg.NewDataType("http://schema.org", "TireShop")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
