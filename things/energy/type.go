package energy

import "github.com/dpb587/go-schemaorg"

// // Properties that take Energy as values are of the form '&lt;Number&gt;
// &lt;Energy unit of measure&gt;'.
var Type = schemaorg.NewDataType("http://schema.org", "Energy")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
