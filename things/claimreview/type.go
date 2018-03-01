package claimreview

import "github.com/dpb587/go-schemaorg"

// // A fact-checking review of claims made (or reported) in some creative work
// (referenced via itemReviewed).
var Type = schemaorg.NewDataType("http://schema.org", "ClaimReview")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
