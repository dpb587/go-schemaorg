package clip

import "github.com/dpb587/go-schemaorg"

var (
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The series to which this episode or season belongs.
	PartOfSeries = schemaorg.NewProperty("partOfSeries")

	// The episode to which this clip belongs.
	PartOfEpisode = schemaorg.NewProperty("partOfEpisode")

	// The season to which this episode belongs.
	PartOfSeason = schemaorg.NewProperty("partOfSeason")

	// The composer of the soundtrack.
	MusicBy = schemaorg.NewProperty("musicBy")

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can
	// be associated with individual items or with a series, episode, clip.
	Directors = schemaorg.NewProperty("directors")

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an
	// event. Directors can be associated with individual items or with a series,
	// episode, clip.
	Director = schemaorg.NewProperty("director")

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated
	// with individual items or with a series, episode, clip.
	Actors = schemaorg.NewProperty("actors")

	// Position of the clip within an ordered group of clips.
	ClipNumber = schemaorg.NewProperty("clipNumber")
)
