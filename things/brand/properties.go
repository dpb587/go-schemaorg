package brand

import "github.com/dpb587/go-schemaorg"

var (
	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating = schemaorg.NewProperty("aggregateRating")

	// An associated logo.
	Logo = schemaorg.NewProperty("logo")

	// A review of the item.
	Review = schemaorg.NewProperty("review")
)
