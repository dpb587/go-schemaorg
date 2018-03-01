package sellaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of participant. The participant/person/organization that
	// bought the object.
	Buyer = schemaorg.NewProperty("buyer")

	// The warranty promise(s) included in the offer.
	WarrantyPromise = schemaorg.NewProperty("warrantyPromise")
)
