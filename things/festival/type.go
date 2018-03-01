package festival

import "github.com/dpb587/go-schemaorg"

// // Event type: Festival.
var Type = schemaorg.NewDataType("http://schema.org", "Festival")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
