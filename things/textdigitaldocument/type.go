package textdigitaldocument

import "github.com/dpb587/go-schemaorg"

// // A file composed primarily of text.
var Type = schemaorg.NewDataType("http://schema.org", "TextDigitalDocument")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
