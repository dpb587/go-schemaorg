package courthouse

import "github.com/dpb587/go-schemaorg"

// // A courthouse.
var Type = schemaorg.NewDataType("http://schema.org", "Courthouse")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
