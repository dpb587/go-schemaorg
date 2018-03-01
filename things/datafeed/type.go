package datafeed

import "github.com/dpb587/go-schemaorg"

// // A single feed providing structured information about one or more entities or
// topics.
var Type = schemaorg.NewDataType("http://schema.org", "DataFeed")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
