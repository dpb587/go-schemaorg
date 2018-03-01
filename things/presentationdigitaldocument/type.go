package presentationdigitaldocument

import "github.com/dpb587/go-schemaorg"

// // A file containing slides or used for a presentation.
var Type = schemaorg.NewDataType("http://schema.org", "PresentationDigitalDocument")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
