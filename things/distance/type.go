package distance

import "github.com/dpb587/go-schemaorg"

// // Properties that take Distances as values are of the form '&lt;Number&gt;
// &lt;Length unit of measure&gt;'. E.g., '7 ft'.
var Type = schemaorg.NewDataType("http://schema.org", "Distance")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
