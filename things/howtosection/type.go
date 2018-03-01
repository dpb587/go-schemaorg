package howtosection

import "github.com/dpb587/go-schemaorg"

// // A sub-grouping of steps in the instructions for how to achieve a result (e.g.
// steps for making a pie crust within a pie recipe).
var Type = schemaorg.NewDataType("http://schema.org", "HowToSection")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
