package imageobject

import "github.com/dpb587/go-schemaorg"

// // An image file.
var Type = schemaorg.NewDataType("http://schema.org", "ImageObject")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
