package clip

import "github.com/dpb587/go-schemaorg"

// // A short TV or radio program or a segment/part of a program.
var Type = schemaorg.NewDataType("http://schema.org", "Clip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
