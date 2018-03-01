package musicalbum

import "github.com/dpb587/go-schemaorg"

// // A collection of music tracks.
var Type = schemaorg.NewDataType("http://schema.org", "MusicAlbum")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
