package itemlist

import "github.com/dpb587/go-schemaorg"

// // A list of items of any sort&#x2014;for example, Top 10 Movies About
// Weathermen, or Top 100 Party Songs. Not to be confused with HTML lists, which
// are often used only for formatting.
var Type = schemaorg.NewDataType("http://schema.org", "ItemList")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
