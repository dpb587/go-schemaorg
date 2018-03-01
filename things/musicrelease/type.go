package musicrelease

import "github.com/dpb587/go-schemaorg"

// // A MusicRelease is a specific release of a music album.
var Type = schemaorg.NewDataType("http://schema.org", "MusicRelease")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
