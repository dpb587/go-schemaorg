package aggregateoffer

import "github.com/dpb587/go-schemaorg"

// // When a single product is associated with multiple offers (for example, the
// same pair of shoes is offered by different merchants), then AggregateOffer
// can be used.
var Type = schemaorg.NewDataType("http://schema.org", "AggregateOffer")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
