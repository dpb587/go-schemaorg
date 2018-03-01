package conversation

import "github.com/dpb587/go-schemaorg"

// // One or more messages between organizations or people on a particular topic.
// Individual messages can be linked to the conversation with isPartOf or
// hasPart properties.
var Type = schemaorg.NewDataType("http://schema.org", "Conversation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
