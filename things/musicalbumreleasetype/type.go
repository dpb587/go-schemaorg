package musicalbumreleasetype

import "github.com/dpb587/go-schemaorg"

// // The kind of release which this album is: single, EP or album.
var Type = schemaorg.NewDataType("http://schema.org", "MusicAlbumReleaseType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
