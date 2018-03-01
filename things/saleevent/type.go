package saleevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Sales event.
var Type = schemaorg.NewDataType("http://schema.org", "SaleEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
