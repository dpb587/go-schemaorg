package musicvideoobject

import "github.com/dpb587/go-schemaorg"

// // A music video file.
var Type = schemaorg.NewDataType("http://schema.org", "MusicVideoObject")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
