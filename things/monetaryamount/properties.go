package monetaryamount

import "github.com/dpb587/go-schemaorg"

var (
	// The lower value of some characteristic or property.
	MinValue = schemaorg.NewProperty("minValue")

	// The currency in which the monetary amount is expressed (in 3-letter <a
	// href="http://en.wikipedia.org/wiki/ISO_4217">ISO 4217</a> format).
	Currency = schemaorg.NewProperty("currency")

	// <p>The value of the quantitative value or property value node.</p>
	// 
	// <ul>
	// <li>For <a class="localLink"
	// href="http://schema.org/QuantitativeValue">QuantitativeValue</a> and <a
	// class="localLink" href="http://schema.org/MonetaryAmount">MonetaryAmount</a>,
	// the recommended type for values is 'Number'.</li>
	// <li>For <a class="localLink"
	// href="http://schema.org/PropertyValue">PropertyValue</a>, it can be 'Text;',
	// 'Number', 'Boolean', or 'StructuredValue'.</li>
	// </ul>
	// 
	Value = schemaorg.NewProperty("value")

	// The date when the item becomes valid.
	ValidFrom = schemaorg.NewProperty("validFrom")

	// The date after when the item is not valid. For example the end of an offer,
	// salary period, or a period of opening hours.
	ValidThrough = schemaorg.NewProperty("validThrough")

	// The upper value of some characteristic or property.
	MaxValue = schemaorg.NewProperty("maxValue")
)
