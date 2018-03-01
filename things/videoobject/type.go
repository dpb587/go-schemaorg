package videoobject

import "github.com/dpb587/go-schemaorg"

// // A video file.
var Type = schemaorg.NewDataType("http://schema.org", "VideoObject")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
