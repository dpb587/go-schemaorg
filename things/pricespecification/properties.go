package pricespecification

import "github.com/dpb587/go-schemaorg"

var (
	// The lowest price if the price is a range.
	MinPrice = schemaorg.NewProperty("minPrice")

	// The transaction volume, in a monetary unit, for which the offer or price
	// specification is valid, e.g. for indicating a minimal purchasing volume, to
	// express free shipping above a certain order volume, or to limit the
	// acceptance of credit cards to purchases to a certain minimal amount.
	EligibleTransactionVolume = schemaorg.NewProperty("eligibleTransactionVolume")

	// The highest price if the price is a range.
	MaxPrice = schemaorg.NewProperty("maxPrice")

	// The currency (in 3-letter ISO 4217 format) of the price or a price component,
	// when attached to <a class="localLink"
	// href="http://schema.org/PriceSpecification">PriceSpecification</a> and its
	// subtypes.
	PriceCurrency = schemaorg.NewProperty("priceCurrency")

	// The interval and unit of measurement of ordering quantities for which the
	// offer or price specification is valid. This allows e.g. specifying that a
	// certain freight charge is valid only for a certain quantity.
	EligibleQuantity = schemaorg.NewProperty("eligibleQuantity")

	// The date when the item becomes valid.
	ValidFrom = schemaorg.NewProperty("validFrom")

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

	// Specifies whether the applicable value-added tax (VAT) is included in the
	// price specification or not.
	ValueAddedTaxIncluded = schemaorg.NewProperty("valueAddedTaxIncluded")
)
