package musicplaylist

import "github.com/dpb587/go-schemaorg"

// // A collection of music tracks in playlist form.
var Type = schemaorg.NewDataType("http://schema.org", "MusicPlaylist")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
