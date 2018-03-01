package datafeeditem

import "github.com/dpb587/go-schemaorg"

// // A single item within a larger data feed.
var Type = schemaorg.NewDataType("http://schema.org", "DataFeedItem")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
