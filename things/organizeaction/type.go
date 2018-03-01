package organizeaction

import "github.com/dpb587/go-schemaorg"

// // The act of manipulating/administering/supervising/controlling one or more
// objects.
var Type = schemaorg.NewDataType("http://schema.org", "OrganizeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
