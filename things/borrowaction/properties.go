package borrowaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of participant. The person that lends the object being
	// borrowed.
	Lender = schemaorg.NewProperty("lender")
)
