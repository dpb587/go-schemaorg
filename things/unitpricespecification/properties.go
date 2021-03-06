package unitpricespecification

import "github.com/dpb587/go-schemaorg"

var (
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters)
	// or a URL. Other codes than the UN/CEFACT Common Code may be used with a
	// prefix followed by a colon.
	UnitCode = schemaorg.NewProperty("unitCode")

	// This property specifies the minimal quantity and rounding increment that will
	// be the basis for the billing. The unit of measurement is specified by the
	// unitCode property.
	BillingIncrement = schemaorg.NewProperty("billingIncrement")

	// The reference quantity for which a certain price applies, e.g. 1 EUR per 4
	// kWh of electricity. This property is a replacement for unitOfMeasurement for
	// the advanced cases where the price does not relate to a standard unit.
	ReferenceQuantity = schemaorg.NewProperty("referenceQuantity")

	// A string or text indicating the unit of measurement. Useful if you cannot
	// provide a standard unit code for
	// <a href='unitCode'>unitCode</a>.
	UnitText = schemaorg.NewProperty("unitText")

	// A short text or acronym indicating multiple price specifications for the same
	// offer, e.g. SRP for the suggested retail price or INVOICE for the invoice
	// price, mostly used in the car industry.
	PriceType = schemaorg.NewProperty("priceType")
)
