package governmentservice

import "github.com/dpb587/go-schemaorg"

var (
	// The operating organization, if different from the provider.  This enables the
	// representation of services that are provided by an organization, but operated
	// by another organization like a subcontractor.
	ServiceOperator = schemaorg.NewProperty("serviceOperator")
)
