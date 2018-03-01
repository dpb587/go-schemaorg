package emailmessage

import "github.com/dpb587/go-schemaorg"

// // An email message.
var Type = schemaorg.NewDataType("http://schema.org", "EmailMessage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
