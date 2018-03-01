package automatedteller

import "github.com/dpb587/go-schemaorg"

// // ATM/cash machine.
var Type = schemaorg.NewDataType("http://schema.org", "AutomatedTeller")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
