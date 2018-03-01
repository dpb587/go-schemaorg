package itemlistordertype

import "github.com/dpb587/go-schemaorg"

// // Enumerated for values for itemListOrder for indicating how an ordered
// ItemList is organized.
var Type = schemaorg.NewDataType("http://schema.org", "ItemListOrderType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
