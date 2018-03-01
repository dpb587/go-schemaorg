package propertyvaluespecification

import "github.com/dpb587/go-schemaorg"

// // A Property value specification.
var Type = schemaorg.NewDataType("http://schema.org", "PropertyValueSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
