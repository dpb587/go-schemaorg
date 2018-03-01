package mass

import "github.com/dpb587/go-schemaorg"

// // Properties that take Mass as values are of the form '&lt;Number&gt; &lt;Mass
// unit of measure&gt;'. E.g., '7 kg'.
var Type = schemaorg.NewDataType("http://schema.org", "Mass")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
