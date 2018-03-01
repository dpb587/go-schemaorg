package paintaction

import "github.com/dpb587/go-schemaorg"

// // The act of producing a painting, typically with paint and canvas as
// instruments.
var Type = schemaorg.NewDataType("http://schema.org", "PaintAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
