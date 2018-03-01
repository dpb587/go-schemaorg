package videogallery

import "github.com/dpb587/go-schemaorg"

// // Web page type: Video gallery page.
var Type = schemaorg.NewDataType("http://schema.org", "VideoGallery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
