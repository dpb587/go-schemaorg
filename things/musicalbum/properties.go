package musicalbum

import "github.com/dpb587/go-schemaorg"

var (
	// Classification of the album by it's type of content: soundtrack, live album,
	// studio album, etc.
	AlbumProductionType = schemaorg.NewProperty("albumProductionType")

	// The kind of release which this album is: single, EP or album.
	AlbumReleaseType = schemaorg.NewProperty("albumReleaseType")

	// The artist that performed this album or recording.
	ByArtist = schemaorg.NewProperty("byArtist")

	// A release of this album.
	AlbumRelease = schemaorg.NewProperty("albumRelease")
)
