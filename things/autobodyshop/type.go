package autobodyshop

import "github.com/dpb587/go-schemaorg"

// // Auto body shop.
var Type = schemaorg.NewDataType("http://schema.org", "AutoBodyShop")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
