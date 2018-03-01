package sculpture

import "github.com/dpb587/go-schemaorg"

// // A piece of sculpture.
var Type = schemaorg.NewDataType("http://schema.org", "Sculpture")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
