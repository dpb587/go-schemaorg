package musicrelease

import "github.com/dpb587/go-schemaorg"

var (
	// The label that issued the release.
	RecordLabel = schemaorg.NewProperty("recordLabel")

	// The catalog number for the release.
	CatalogNumber = schemaorg.NewProperty("catalogNumber")

	// The group the release is credited to if different than the byArtist. For
	// example, Red and Blue is credited to "Stefani Germanotta Band", but by Lady
	// Gaga.
	CreditedTo = schemaorg.NewProperty("creditedTo")

	// The album this is a release of.
	ReleaseOf = schemaorg.NewProperty("releaseOf")

	// The duration of the item (movie, audio recording, event, etc.) in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>.
	Duration = schemaorg.NewProperty("duration")

	// Format of this release (the type of recording media used, ie. compact disc,
	// digital media, LP, etc.).
	MusicReleaseFormat = schemaorg.NewProperty("musicReleaseFormat")
)
