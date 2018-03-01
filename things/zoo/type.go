package zoo

import "github.com/dpb587/go-schemaorg"

// // A zoo.
var Type = schemaorg.NewDataType("http://schema.org", "Zoo")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
