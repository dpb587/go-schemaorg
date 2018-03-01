package quantitativevalue

import "github.com/dpb587/go-schemaorg"

var (
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters)
	// or a URL. Other codes than the UN/CEFACT Common Code may be used with a
	// prefix followed by a colon.
	UnitCode = schemaorg.NewProperty("unitCode")

	// The lower value of some characteristic or property.
	MinValue = schemaorg.NewProperty("minValue")

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

	// A property-value pair representing an additional characteristics of the
	// entitity, e.g. a product feature or another characteristic for which there is
	// no matching property in schema.org.</p>
	// 
	// <p>Note: Publishers should be aware that applications designed to use
	// specific schema.org properties (e.g. http://schema.org/width,
	// http://schema.org/color, http://schema.org/gtin13, ...) will typically expect
	// such data to be provided using those properties, rather than using the
	// generic property/value mechanism.
	AdditionalProperty = schemaorg.NewProperty("additionalProperty")

	// A pointer to a secondary value that provides additional information on the
	// original value, e.g. a reference temperature.
	ValueReference = schemaorg.NewProperty("valueReference")

	// The upper value of some characteristic or property.
	MaxValue = schemaorg.NewProperty("maxValue")

	// A string or text indicating the unit of measurement. Useful if you cannot
	// provide a standard unit code for
	// <a href='unitCode'>unitCode</a>.
	UnitText = schemaorg.NewProperty("unitText")
)
