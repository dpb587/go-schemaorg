package ondemandevent

import "github.com/dpb587/go-schemaorg"

// // A publication event e.g. catch-up TV or radio podcast, during which a program
// is available on-demand.
var Type = schemaorg.NewDataType("http://schema.org", "OnDemandEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
