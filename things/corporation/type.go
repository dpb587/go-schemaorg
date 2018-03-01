package corporation

import "github.com/dpb587/go-schemaorg"

// // Organization: A business corporation.
var Type = schemaorg.NewDataType("http://schema.org", "Corporation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
