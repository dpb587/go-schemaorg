package playaction

import "github.com/dpb587/go-schemaorg"

var (
	// An intended audience, i.e. a group for whom something was created.
	Audience = schemaorg.NewProperty("audience")

	// Upcoming or past event associated with this place, organization, or action.
	Event = schemaorg.NewProperty("event")
)
