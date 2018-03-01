package digitaldocument

import "github.com/dpb587/go-schemaorg"

// // An electronic file or document.
var Type = schemaorg.NewDataType("http://schema.org", "DigitalDocument")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
