package bookformattype

import "github.com/dpb587/go-schemaorg"

// // The publication format of the book.
var Type = schemaorg.NewDataType("http://schema.org", "BookFormatType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
