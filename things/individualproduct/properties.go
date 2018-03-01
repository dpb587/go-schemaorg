package individualproduct

import "github.com/dpb587/go-schemaorg"

var (
	// The serial number or any alphanumeric identifier of a particular product.
	// When attached to an offer, it is a shortcut for the serial number of the
	// product included in the offer.
	SerialNumber = schemaorg.NewProperty("serialNumber")
)
