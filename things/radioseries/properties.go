package radioseries

import "github.com/dpb587/go-schemaorg"

var (
	// A season in a media series.
	Seasons = schemaorg.NewProperty("seasons")

	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The number of episodes in this season or series.
	NumberOfEpisodes = schemaorg.NewProperty("numberOfEpisodes")

	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer = schemaorg.NewProperty("trailer")

	// A season in a media series.
	Season = schemaorg.NewProperty("season")

	// An episode of a TV/radio series or season.
	Episodes = schemaorg.NewProperty("episodes")

	// The composer of the soundtrack.
	MusicBy = schemaorg.NewProperty("musicBy")

	// A season that is part of the media series.
	ContainsSeason = schemaorg.NewProperty("containsSeason")

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can
	// be associated with individual items or with a series, episode, clip.
	Directors = schemaorg.NewProperty("directors")

	// The number of seasons in this series.
	NumberOfSeasons = schemaorg.NewProperty("numberOfSeasons")

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

	// An episode of a tv, radio or game media within a series or season.
	Episode = schemaorg.NewProperty("episode")
)
