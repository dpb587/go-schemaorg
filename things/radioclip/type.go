package radioclip

import "github.com/dpb587/go-schemaorg"

// // A short radio program or a segment/part of a radio program.
var Type = schemaorg.NewDataType("http://schema.org", "RadioClip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
