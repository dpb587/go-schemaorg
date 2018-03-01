package deactivateaction

import "github.com/dpb587/go-schemaorg"

// // The act of stopping or deactivating a device or application (e.g. stopping a
// timer or turning off a flashlight).
var Type = schemaorg.NewDataType("http://schema.org", "DeactivateAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
