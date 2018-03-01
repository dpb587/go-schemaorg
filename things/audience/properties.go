package audience

import "github.com/dpb587/go-schemaorg"

var (
	// The target group associated with a given audience (e.g. veterans, car owners,
	// musicians, etc.).
	AudienceType = schemaorg.NewProperty("audienceType")

	// The geographic area associated with the audience.
	GeographicArea = schemaorg.NewProperty("geographicArea")
)
