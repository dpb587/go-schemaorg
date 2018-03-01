package movieseries

import "github.com/dpb587/go-schemaorg"

var (
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer = schemaorg.NewProperty("trailer")

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

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated
	// with individual items or with a series, episode, clip.
	Actors = schemaorg.NewProperty("actors")
)
