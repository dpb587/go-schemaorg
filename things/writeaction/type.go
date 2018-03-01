package writeaction

import "github.com/dpb587/go-schemaorg"

// // The act of authoring written creative content.
var Type = schemaorg.NewDataType("http://schema.org", "WriteAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
