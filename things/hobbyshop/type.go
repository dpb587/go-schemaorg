package hobbyshop

import "github.com/dpb587/go-schemaorg"

// // A store that sells materials useful or necessary for various hobbies.
var Type = schemaorg.NewDataType("http://schema.org", "HobbyShop")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
