package loseaction

import "github.com/dpb587/go-schemaorg"

// // The act of being defeated in a competitive activity.
var Type = schemaorg.NewDataType("http://schema.org", "LoseAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
