package drawaction

import "github.com/dpb587/go-schemaorg"

// // The act of producing a visual/graphical representation of an object,
// typically with a pen/pencil and paper as instruments.
var Type = schemaorg.NewDataType("http://schema.org", "DrawAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
