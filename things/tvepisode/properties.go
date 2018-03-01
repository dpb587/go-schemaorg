package tvepisode

import "github.com/dpb587/go-schemaorg"

var (
	// The TV series to which this episode or season belongs.
	PartOfTVSeries = schemaorg.NewProperty("partOfTVSeries")

	// Languages in which subtitles/captions are available, in <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard format</a>.
	SubtitleLanguage = schemaorg.NewProperty("subtitleLanguage")

	// The country of the principal offices of the production company or individual
	// responsible for the movie or program.
	CountryOfOrigin = schemaorg.NewProperty("countryOfOrigin")
)
