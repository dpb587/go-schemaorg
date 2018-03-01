package aggregaterating

import "github.com/dpb587/go-schemaorg"

var (
	// The item that is being reviewed/rated.
	ItemReviewed = schemaorg.NewProperty("itemReviewed")

	// The count of total number of reviews.
	ReviewCount = schemaorg.NewProperty("reviewCount")

	// The count of total number of ratings.
	RatingCount = schemaorg.NewProperty("ratingCount")
)
