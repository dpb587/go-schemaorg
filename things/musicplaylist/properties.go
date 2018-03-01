package musicplaylist

import "github.com/dpb587/go-schemaorg"

var (
	// A music recording (track)&#x2014;usually a single song. If an ItemList is
	// given, the list should contain items of type MusicRecording.
	Track = schemaorg.NewProperty("track")

	// The number of tracks in this album or playlist.
	NumTracks = schemaorg.NewProperty("numTracks")

	// A music recording (track)&#x2014;usually a single song.
	Tracks = schemaorg.NewProperty("tracks")
)
