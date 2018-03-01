package permit

import "github.com/dpb587/go-schemaorg"

var (
	// The time validity of the permit.
	ValidFor = schemaorg.NewProperty("validFor")

	// The date when the item is no longer valid.
	ValidUntil = schemaorg.NewProperty("validUntil")

	// The date when the item becomes valid.
	ValidFrom = schemaorg.NewProperty("validFrom")

	// The organization issuing the ticket or permit.
	IssuedBy = schemaorg.NewProperty("issuedBy")

	// The service through with the permit was granted.
	IssuedThrough = schemaorg.NewProperty("issuedThrough")

	// The target audience for this permit.
	PermitAudience = schemaorg.NewProperty("permitAudience")

	// The geographic area where the permit is valid.
	ValidIn = schemaorg.NewProperty("validIn")
)
