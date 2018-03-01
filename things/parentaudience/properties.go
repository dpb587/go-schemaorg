package parentaudience

import "github.com/dpb587/go-schemaorg"

var (
	// Maximal age of the child.
	ChildMaxAge = schemaorg.NewProperty("childMaxAge")

	// Minimal age of the child.
	ChildMinAge = schemaorg.NewProperty("childMinAge")
)
