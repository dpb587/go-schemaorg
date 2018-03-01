package review

import "github.com/dpb587/go-schemaorg"

// // A review of an item - for example, of a restaurant, movie, or store.
var Type = schemaorg.NewDataType("http://schema.org", "Review")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
