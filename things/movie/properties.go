package movie

import "github.com/dpb587/go-schemaorg"

var (
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer = schemaorg.NewProperty("trailer")

	// Languages in which subtitles/captions are available, in <a
	// href="http://tools.ietf.org/html/bcp47">IETF BCP 47 standard format</a>.
	SubtitleLanguage = schemaorg.NewProperty("subtitleLanguage")

	// The country of the principal offices of the production company or individual
	// responsible for the movie or program.
	CountryOfOrigin = schemaorg.NewProperty("countryOfOrigin")

	// The composer of the soundtrack.
	MusicBy = schemaorg.NewProperty("musicBy")

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can
	// be associated with individual items or with a series, episode, clip.
	Directors = schemaorg.NewProperty("directors")

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an
	// event. Directors can be associated with individual items or with a series,
	// episode, clip.
	Director = schemaorg.NewProperty("director")

	// The production company or studio responsible for the item e.g. series, video
	// game, episode etc.
	ProductionCompany = schemaorg.NewProperty("productionCompany")

	// The duration of the item (movie, audio recording, event, etc.) in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>.
	Duration = schemaorg.NewProperty("duration")

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated
	// with individual items or with a series, episode, clip.
	Actors = schemaorg.NewProperty("actors")
)
