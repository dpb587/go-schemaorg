package musicgroup

import "github.com/dpb587/go-schemaorg"

var (
	// A collection of music albums.
	Albums = schemaorg.NewProperty("albums")

	// A music recording (track)&#x2014;usually a single song. If an ItemList is
	// given, the list should contain items of type MusicRecording.
	Track = schemaorg.NewProperty("track")

	// Genre of the creative work, broadcast channel or group.
	Genre = schemaorg.NewProperty("genre")

	// A music recording (track)&#x2014;usually a single song.
	Tracks = schemaorg.NewProperty("tracks")

	// A member of a music group&#x2014;for example, John, Paul, George, or Ringo.
	MusicGroupMember = schemaorg.NewProperty("musicGroupMember")

	// A music album.
	Album = schemaorg.NewProperty("album")
)
