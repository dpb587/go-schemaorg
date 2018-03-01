package residence

import "github.com/dpb587/go-schemaorg"

// // The place where a person lives.
var Type = schemaorg.NewDataType("http://schema.org", "Residence")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
