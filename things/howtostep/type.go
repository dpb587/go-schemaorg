package howtostep

import "github.com/dpb587/go-schemaorg"

// // A step in the instructions for how to achieve a result. It is an ordered list
// with HowToDirection and/or HowToTip items.
var Type = schemaorg.NewDataType("http://schema.org", "HowToStep")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
