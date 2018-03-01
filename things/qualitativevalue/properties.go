package qualitativevalue

import "github.com/dpb587/go-schemaorg"

var (
	// This ordering relation for qualitative values indicates that the subject is
	// greater than the object.
	Greater = schemaorg.NewProperty("greater")

	// This ordering relation for qualitative values indicates that the subject is
	// equal to the object.
	Equal = schemaorg.NewProperty("equal")

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

	// This ordering relation for qualitative values indicates that the subject is
	// lesser than the object.
	Lesser = schemaorg.NewProperty("lesser")

	// A pointer to a secondary value that provides additional information on the
	// original value, e.g. a reference temperature.
	ValueReference = schemaorg.NewProperty("valueReference")

	// This ordering relation for qualitative values indicates that the subject is
	// not equal to the object.
	NonEqual = schemaorg.NewProperty("nonEqual")

	// This ordering relation for qualitative values indicates that the subject is
	// lesser than or equal to the object.
	LesserOrEqual = schemaorg.NewProperty("lesserOrEqual")

	// This ordering relation for qualitative values indicates that the subject is
	// greater than or equal to the object.
	GreaterOrEqual = schemaorg.NewProperty("greaterOrEqual")
)
