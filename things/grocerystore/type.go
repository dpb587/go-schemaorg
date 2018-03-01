package grocerystore

import "github.com/dpb587/go-schemaorg"

// // A grocery store.
var Type = schemaorg.NewDataType("http://schema.org", "GroceryStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
