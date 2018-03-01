package activateaction

import "github.com/dpb587/go-schemaorg"

// // The act of starting or activating a device or application (e.g. starting a
// timer or turning on a flashlight).
var Type = schemaorg.NewDataType("http://schema.org", "ActivateAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
