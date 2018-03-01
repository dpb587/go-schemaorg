package product

import "github.com/dpb587/go-schemaorg"

// // Any offered product or service. For example: a pair of shoes; a concert
// ticket; the rental of a car; a haircut; or an episode of a TV show streamed
// online.
var Type = schemaorg.NewDataType("http://schema.org", "Product")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
