package rating

import "github.com/dpb587/go-schemaorg"

var (
	// The rating for the content.
	RatingValue = schemaorg.NewProperty("ratingValue")

	// The highest value allowed in this rating system. If bestRating is omitted, 5
	// is assumed.
	BestRating = schemaorg.NewProperty("bestRating")

	// The author of this content or rating. Please note that author is special in
	// that HTML 5 provides a special mechanism for indicating authorship via the
	// rel tag. That is equivalent to this and may be used interchangeably.
	Author = schemaorg.NewProperty("author")

	// The lowest value allowed in this rating system. If worstRating is omitted, 1
	// is assumed.
	WorstRating = schemaorg.NewProperty("worstRating")
)
