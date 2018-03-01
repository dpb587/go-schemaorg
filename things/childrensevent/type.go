package childrensevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Children's event.
var Type = schemaorg.NewDataType("http://schema.org", "ChildrensEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
