package aggregateoffer

import "github.com/dpb587/go-schemaorg"

var (
	// An offer to provide this item&#x2014;for example, an offer to sell a product,
	// rent the DVD of a movie, perform a service, or give away tickets to an event.
	Offers = schemaorg.NewProperty("offers")

	// The number of offers for the product.
	OfferCount = schemaorg.NewProperty("offerCount")

	// The lowest price of all offers available.
	LowPrice = schemaorg.NewProperty("lowPrice")

	// The highest price of all offers available.
	HighPrice = schemaorg.NewProperty("highPrice")
)
