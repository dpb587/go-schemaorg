package message

import "github.com/dpb587/go-schemaorg"

// // A single message from a sender to one or more organizations or people.
var Type = schemaorg.NewDataType("http://schema.org", "Message")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
