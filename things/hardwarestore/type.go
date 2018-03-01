package hardwarestore

import "github.com/dpb587/go-schemaorg"

// // A hardware store.
var Type = schemaorg.NewDataType("http://schema.org", "HardwareStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
