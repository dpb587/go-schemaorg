package howtodirection

import "github.com/dpb587/go-schemaorg"

// // A direction indicating a single action to do in the instructions for how to
// achieve a result.
var Type = schemaorg.NewDataType("http://schema.org", "HowToDirection")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
