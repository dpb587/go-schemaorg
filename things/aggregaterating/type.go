package aggregaterating

import "github.com/dpb587/go-schemaorg"

// // The average rating based on multiple ratings or reviews.
var Type = schemaorg.NewDataType("http://schema.org", "AggregateRating")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
