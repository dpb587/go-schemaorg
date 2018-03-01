package person

import "github.com/dpb587/go-schemaorg"

// // A person (alive, dead, undead, or fictional).
var Type = schemaorg.NewDataType("http://schema.org", "Person")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
