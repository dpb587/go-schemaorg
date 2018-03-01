package firestation

import "github.com/dpb587/go-schemaorg"

// // A fire station. With firemen.
var Type = schemaorg.NewDataType("http://schema.org", "FireStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
