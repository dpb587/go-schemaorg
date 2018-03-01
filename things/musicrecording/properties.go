package musicrecording

import "github.com/dpb587/go-schemaorg"

var (
	// The composition this track is a recording of.
	RecordingOf = schemaorg.NewProperty("recordingOf")

	// The International Standard Recording Code for the recording.
	IsrcCode = schemaorg.NewProperty("isrcCode")

	// The album to which this recording belongs.
	InAlbum = schemaorg.NewProperty("inAlbum")

	// The playlist to which this recording belongs.
	InPlaylist = schemaorg.NewProperty("inPlaylist")

	// The artist that performed this album or recording.
	ByArtist = schemaorg.NewProperty("byArtist")

	// The duration of the item (movie, audio recording, event, etc.) in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>.
	Duration = schemaorg.NewProperty("duration")
)
