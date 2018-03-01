package deliverychargespecification

import "github.com/dpb587/go-schemaorg"

var (
	// The geographic area where a service or offered item is provided.
	AreaServed = schemaorg.NewProperty("areaServed")

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the
	// GeoShape for the geo-political region(s) for which the offer or delivery
	// charge specification is not valid, e.g. a region where the transaction is not
	// allowed.</p>
	// 
	// <p>See also <a class="localLink"
	// href="http://schema.org/eligibleRegion">eligibleRegion</a>.
	IneligibleRegion = schemaorg.NewProperty("ineligibleRegion")

	// The delivery method(s) to which the delivery charge or payment charge
	// specification applies.
	AppliesToDeliveryMethod = schemaorg.NewProperty("appliesToDeliveryMethod")

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the
	// GeoShape for the geo-political region(s) for which the offer or delivery
	// charge specification is valid.</p>
	// 
	// <p>See also <a class="localLink"
	// href="http://schema.org/ineligibleRegion">ineligibleRegion</a>.
	EligibleRegion = schemaorg.NewProperty("eligibleRegion")
)
