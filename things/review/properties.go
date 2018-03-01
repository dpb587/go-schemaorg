package review

import "github.com/dpb587/go-schemaorg"

var (
	// The rating given in this review. Note that reviews can themselves be rated.
	// The <code>reviewRating</code> applies to rating given by the review. The <a
	// class="localLink"
	// href="http://schema.org/aggregateRating">aggregateRating</a> property applies
	// to the review itself, as a creative work.
	ReviewRating = schemaorg.NewProperty("reviewRating")

	// The item that is being reviewed/rated.
	ItemReviewed = schemaorg.NewProperty("itemReviewed")

	// The actual body of the review.
	ReviewBody = schemaorg.NewProperty("reviewBody")
)
