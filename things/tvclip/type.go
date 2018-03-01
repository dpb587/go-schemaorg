package tvclip

import "github.com/dpb587/go-schemaorg"

// // A short TV program or a segment/part of a TV program.
var Type = schemaorg.NewDataType("http://schema.org", "TVClip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
