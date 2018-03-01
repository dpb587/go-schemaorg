package locationfeaturespecification

import "github.com/dpb587/go-schemaorg"

// // Specifies a location feature by providing a structured value representing a
// feature of an accommodation as a property-value pair of varying degrees of
// formality.
var Type = schemaorg.NewDataType("http://schema.org", "LocationFeatureSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
