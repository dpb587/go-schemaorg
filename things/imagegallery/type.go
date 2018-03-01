package imagegallery

import "github.com/dpb587/go-schemaorg"

// // Web page type: Image gallery page.
var Type = schemaorg.NewDataType("http://schema.org", "ImageGallery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
