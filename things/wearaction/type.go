package wearaction

import "github.com/dpb587/go-schemaorg"

// // The act of dressing oneself in clothing.
var Type = schemaorg.NewDataType("http://schema.org", "WearAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
