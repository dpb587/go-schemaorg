package parceldelivery

import "github.com/dpb587/go-schemaorg"

var (
	// The service provider, service operator, or service performer; the goods
	// producer. Another party (a seller) may offer those services or goods on
	// behalf of the provider. A provider may also serve as the seller.
	Provider = schemaorg.NewProperty("provider")

	// Tracking url for the parcel delivery.
	TrackingUrl = schemaorg.NewProperty("trackingUrl")

	// Destination address.
	DeliveryAddress = schemaorg.NewProperty("deliveryAddress")

	// Shipper tracking number.
	TrackingNumber = schemaorg.NewProperty("trackingNumber")

	// Method used for delivery or shipping.
	HasDeliveryMethod = schemaorg.NewProperty("hasDeliveryMethod")

	// New entry added as the package passes through each leg of its journey (from
	// shipment to final delivery).
	DeliveryStatus = schemaorg.NewProperty("deliveryStatus")

	// 'carrier' is an out-dated term indicating the 'provider' for parcel delivery
	// and flights.
	Carrier = schemaorg.NewProperty("carrier")

	// Shipper's address.
	OriginAddress = schemaorg.NewProperty("originAddress")

	// Item(s) being shipped.
	ItemShipped = schemaorg.NewProperty("itemShipped")

	// The overall order the items in this delivery were included in.
	PartOfOrder = schemaorg.NewProperty("partOfOrder")

	// The earliest date the package may arrive.
	ExpectedArrivalFrom = schemaorg.NewProperty("expectedArrivalFrom")

	// The latest date the package may arrive.
	ExpectedArrivalUntil = schemaorg.NewProperty("expectedArrivalUntil")
)
