package itempage

import "github.com/dpb587/go-schemaorg"

// // A page devoted to a single item, such as a particular product or hotel.
var Type = schemaorg.NewDataType("http://schema.org", "ItemPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
