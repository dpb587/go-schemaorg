package controlaction

import "github.com/dpb587/go-schemaorg"

// // An agent controls a device or application.
var Type = schemaorg.NewDataType("http://schema.org", "ControlAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
