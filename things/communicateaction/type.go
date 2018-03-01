package communicateaction

import "github.com/dpb587/go-schemaorg"

// // The act of conveying information to another person via a communication medium
// (instrument) such as speech, email, or telephone conversation.
var Type = schemaorg.NewDataType("http://schema.org", "CommunicateAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
