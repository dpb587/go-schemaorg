package park

import "github.com/dpb587/go-schemaorg"

// // A park.
var Type = schemaorg.NewDataType("http://schema.org", "Park")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
