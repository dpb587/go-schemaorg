package artgallery

import "github.com/dpb587/go-schemaorg"

// // An art gallery.
var Type = schemaorg.NewDataType("http://schema.org", "ArtGallery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
