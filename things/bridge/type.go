package bridge

import "github.com/dpb587/go-schemaorg"

// // A bridge.
var Type = schemaorg.NewDataType("http://schema.org", "Bridge")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
