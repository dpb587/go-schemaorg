package notedigitaldocument

import "github.com/dpb587/go-schemaorg"

// // A file containing a note, primarily for the author.
var Type = schemaorg.NewDataType("http://schema.org", "NoteDigitalDocument")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
