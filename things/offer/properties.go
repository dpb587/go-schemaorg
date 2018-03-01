package offer

import "github.com/dpb587/go-schemaorg"

var (
	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a
	// product or service, or the product to which the offer refers.
	Sku = schemaorg.NewProperty("sku")

	// The beginning of the availability of the product or service included in the
	// offer.
	AvailabilityStarts = schemaorg.NewProperty("availabilityStarts")

	// The delivery method(s) available for this offer.
	AvailableDeliveryMethod = schemaorg.NewProperty("availableDeliveryMethod")

	// The geographic area where a service or offered item is provided.
	AreaServed = schemaorg.NewProperty("areaServed")

	// The Manufacturer Part Number (MPN) of the product, or the product to which
	// the offer refers.
	Mpn = schemaorg.NewProperty("mpn")

	// The serial number or any alphanumeric identifier of a particular product.
	// When attached to an offer, it is a shortcut for the serial number of the
	// product included in the offer.
	SerialNumber = schemaorg.NewProperty("serialNumber")

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the
	// GeoShape for the geo-political region(s) for which the offer or delivery
	// charge specification is not valid, e.g. a region where the transaction is not
	// allowed.</p>
	// 
	// <p>See also <a class="localLink"
	// href="http://schema.org/eligibleRegion">eligibleRegion</a>.
	IneligibleRegion = schemaorg.NewProperty("ineligibleRegion")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx">GTIN-8</a>
	// code of the product, or the product to which the offer refers. This code is
	// also known as EAN/UCC-8 or 8-digit EAN. See <a
	// href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN Summary</a>
	// for more details.
	Gtin8 = schemaorg.NewProperty("gtin8")

	// Review of the item.
	Reviews = schemaorg.NewProperty("reviews")

	// The availability of this item&#x2014;for example In stock, Out of stock,
	// Pre-order, etc.
	Availability = schemaorg.NewProperty("availability")

	// One or more detailed price specifications, indicating the unit price and
	// delivery or payment charges.
	PriceSpecification = schemaorg.NewProperty("priceSpecification")

	// The current approximate inventory level for the item or items.
	InventoryLevel = schemaorg.NewProperty("inventoryLevel")

	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating = schemaorg.NewProperty("aggregateRating")

	// The transaction volume, in a monetary unit, for which the offer or price
	// specification is valid, e.g. for indicating a minimal purchasing volume, to
	// express free shipping above a certain order volume, or to limit the
	// acceptance of credit cards to purchases to a certain minimal amount.
	EligibleTransactionVolume = schemaorg.NewProperty("eligibleTransactionVolume")

	// The place(s) from which the offer can be obtained (e.g. store locations).
	AvailableAtOrFrom = schemaorg.NewProperty("availableAtOrFrom")

	// The currency (in 3-letter ISO 4217 format) of the price or a price component,
	// when attached to <a class="localLink"
	// href="http://schema.org/PriceSpecification">PriceSpecification</a> and its
	// subtypes.
	PriceCurrency = schemaorg.NewProperty("priceCurrency")

	// A category for the item. Greater signs or slashes can be used to informally
	// indicate a category hierarchy.
	Category = schemaorg.NewProperty("category")

	// The date after which the price is no longer available.
	PriceValidUntil = schemaorg.NewProperty("priceValidUntil")

	// An additional offer that can only be obtained in combination with the first
	// base offer (e.g. supplements and extensions that are available for a
	// surcharge).
	AddOn = schemaorg.NewProperty("addOn")

	// The warranty promise(s) included in the offer.
	Warranty = schemaorg.NewProperty("warranty")

	// This links to a node or nodes indicating the exact quantity of the products
	// included in the offer.
	IncludesObject = schemaorg.NewProperty("includesObject")

	// The interval and unit of measurement of ordering quantities for which the
	// offer or price specification is valid. This allows e.g. specifying that a
	// certain freight charge is valid only for a certain quantity.
	EligibleQuantity = schemaorg.NewProperty("eligibleQuantity")

	// The date when the item becomes valid.
	ValidFrom = schemaorg.NewProperty("validFrom")

	// The amount of time that is required between accepting the offer and the
	// actual usage of the resource or service.
	AdvanceBookingRequirement = schemaorg.NewProperty("advanceBookingRequirement")

	// The date after when the item is not valid. For example the end of an offer,
	// salary period, or a period of opening hours.
	ValidThrough = schemaorg.NewProperty("validThrough")

	// <p>The offer price of a product, or of a price component when attached to
	// PriceSpecification and its subtypes.</p>
	// 
	// <p>Usage guidelines:</p>
	// 
	// <ul>
	// <li>Use the <a class="localLink"
	// href="http://schema.org/priceCurrency">priceCurrency</a> property (with <a
	// href="http://en.wikipedia.org/wiki/ISO_4217#Active_codes">ISO 4217 codes</a>
	// e.g. "USD") instead of
	//   including <a
	// href="http://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign">ambiguous
	// symbols</a> such as '$' in the value.</li>
	// <li>Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a
	// decimal point. Avoid using these symbols as a readability separator.</li>
	// <li>Note that both <a
	// href="http://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute">RDFa</a>
	// and Microdata syntax allow the use of a "content=" attribute for publishing
	// simple machine-readable values alongside more human-friendly formatting.</li>
	// <li>Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE'
	// (U+0039)) rather than superficially similiar Unicode symbols.</li>
	// </ul>
	// 
	Price = schemaorg.NewProperty("price")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx">GTIN-14</a>
	// code of the product, or the product to which the offer refers. See <a
	// href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN Summary</a>
	// for more details.
	Gtin14 = schemaorg.NewProperty("gtin14")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx">GTIN-13</a>
	// code of the product, or the product to which the offer refers. This is
	// equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes
	// can be converted into a GTIN-13 code by simply adding a preceeding zero. See
	// <a href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN
	// Summary</a> for more details.
	Gtin13 = schemaorg.NewProperty("gtin13")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx">GTIN-12</a>
	// code of the product, or the product to which the offer refers. The GTIN-12 is
	// the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item
	// Reference, and Check Digit used to identify trade items. See <a
	// href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN Summary</a>
	// for more details.
	Gtin12 = schemaorg.NewProperty("gtin12")

	// An entity which offers (sells / leases / lends / loans) the services / goods.
	// A seller may also be a provider.
	Seller = schemaorg.NewProperty("seller")

	// The typical delay between the receipt of the order and the goods either
	// leaving the warehouse or being prepared for pickup, in case the delivery
	// method is on site pickup.
	DeliveryLeadTime = schemaorg.NewProperty("deliveryLeadTime")

	// The end of the availability of the product or service included in the offer.
	AvailabilityEnds = schemaorg.NewProperty("availabilityEnds")

	// A review of the item.
	Review = schemaorg.NewProperty("review")

	// The business function (e.g. sell, lease, repair, dispose) of the offer or
	// component of a bundle (TypeAndQuantityNode). The default is
	// http://purl.org/goodrelations/v1#Sell.
	BusinessFunction = schemaorg.NewProperty("businessFunction")

	// The type(s) of customers for which the given offer is valid.
	EligibleCustomerType = schemaorg.NewProperty("eligibleCustomerType")

	// A predefined value from OfferItemCondition or a textual description of the
	// condition of the product or service, or the products or services included in
	// the offer.
	ItemCondition = schemaorg.NewProperty("itemCondition")

	// The item being offered.
	ItemOffered = schemaorg.NewProperty("itemOffered")

	// A pointer to the organization or person making the offer.
	OfferedBy = schemaorg.NewProperty("offeredBy")

	// The duration for which the given offer is valid.
	EligibleDuration = schemaorg.NewProperty("eligibleDuration")

	// The payment method(s) accepted by seller for this offer.
	AcceptedPaymentMethod = schemaorg.NewProperty("acceptedPaymentMethod")

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the
	// GeoShape for the geo-political region(s) for which the offer or delivery
	// charge specification is valid.</p>
	// 
	// <p>See also <a class="localLink"
	// href="http://schema.org/ineligibleRegion">ineligibleRegion</a>.
	EligibleRegion = schemaorg.NewProperty("eligibleRegion")
)
