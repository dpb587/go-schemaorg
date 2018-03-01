package consumeaction

import "github.com/dpb587/go-schemaorg"

var (
	// An Offer which must be accepted before the user can perform the Action. For
	// example, the user may need to buy a movie before being able to watch it.
	ExpectsAcceptanceOf = schemaorg.NewProperty("expectsAcceptanceOf")
)
