package exhibitionevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Exhibition event, e.g. at a museum, library, archive, tradeshow,
// ...
var Type = schemaorg.NewDataType("http://schema.org", "ExhibitionEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
