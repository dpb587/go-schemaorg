package structuredvalue

import "github.com/dpb587/go-schemaorg"

// // Structured values are used when the value of a property has a more complex
// structure than simply being a textual value or a reference to another thing.
var Type = schemaorg.NewDataType("http://schema.org", "StructuredValue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
