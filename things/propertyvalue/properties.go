package propertyvalue

import "github.com/dpb587/go-schemaorg"

var (
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters)
	// or a URL. Other codes than the UN/CEFACT Common Code may be used with a
	// prefix followed by a colon.
	UnitCode = schemaorg.NewProperty("unitCode")

	// The lower value of some characteristic or property.
	MinValue = schemaorg.NewProperty("minValue")

	// A commonly used identifier for the characteristic represented by the
	// property, e.g. a manufacturer or a standard code for a property. propertyID
	// can be
	// (1) a prefixed string, mainly meant to be used with standards for product
	// properties; (2) a site-specific, non-prefixed string (e.g. the primary key of
	// the property or the vendor-specific id of the property), or (3)
	// a URL indicating the type of the property, either pointing to an external
	// vocabulary, or a Web resource that describes the property (e.g. a glossary
	// entry).
	// Standards bodies should promote a standard prefix for the identifiers of
	// properties from their standards.
	PropertyID = schemaorg.NewProperty("propertyID")

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
