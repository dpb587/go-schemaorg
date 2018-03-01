package ownershipinfo

import "github.com/dpb587/go-schemaorg"

var (
	// The product that this structured value is referring to.
	TypeOfGood = schemaorg.NewProperty("typeOfGood")

	// The date and time of giving up ownership on the product.
	OwnedThrough = schemaorg.NewProperty("ownedThrough")

	// The organization or person from which the product was acquired.
	AcquiredFrom = schemaorg.NewProperty("acquiredFrom")

	// The date and time of obtaining the product.
	OwnedFrom = schemaorg.NewProperty("ownedFrom")
)
