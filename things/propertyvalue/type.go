package propertyvalue

import "github.com/dpb587/go-schemaorg"

// // A property-value pair, e.g. representing a feature of a product or place. Use
// the 'name' property for the name of the property. If there is an additional
// human-readable version of the value, put that into the 'description'
// property.</p>
// 
// <p>Always use specific schema.org properties when a) they exist and b) you
// can populate them. Using PropertyValue as a substitute will typically not
// trigger the same effect as using the original, specific property.
var Type = schemaorg.NewDataType("http://schema.org", "PropertyValue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
