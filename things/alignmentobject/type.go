package alignmentobject

import "github.com/dpb587/go-schemaorg"

// // An intangible item that describes an alignment between a learning resource
// and a node in an educational framework.
var Type = schemaorg.NewDataType("http://schema.org", "AlignmentObject")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
