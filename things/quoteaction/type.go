package quoteaction

import "github.com/dpb587/go-schemaorg"

// // An agent quotes/estimates/appraises an object/product/service with a price at
// a location/store.
var Type = schemaorg.NewDataType("http://schema.org", "QuoteAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
