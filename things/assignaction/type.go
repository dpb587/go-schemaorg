package assignaction

import "github.com/dpb587/go-schemaorg"

// // The act of allocating an action/event/task to some destination (someone or
// something).
var Type = schemaorg.NewDataType("http://schema.org", "AssignAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
