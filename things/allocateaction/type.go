package allocateaction

import "github.com/dpb587/go-schemaorg"

// // The act of organizing tasks/objects/events by associating resources to it.
var Type = schemaorg.NewDataType("http://schema.org", "AllocateAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
