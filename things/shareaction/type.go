package shareaction

import "github.com/dpb587/go-schemaorg"

// // The act of distributing content to people for their amusement or edification.
var Type = schemaorg.NewDataType("http://schema.org", "ShareAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
