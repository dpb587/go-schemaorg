package creativeworkseason

import "github.com/dpb587/go-schemaorg"

var (
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The series to which this episode or season belongs.
	PartOfSeries = schemaorg.NewProperty("partOfSeries")

	// The number of episodes in this season or series.
	NumberOfEpisodes = schemaorg.NewProperty("numberOfEpisodes")

	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer = schemaorg.NewProperty("trailer")

	// The end date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	EndDate = schemaorg.NewProperty("endDate")

	// An episode of a TV/radio series or season.
	Episodes = schemaorg.NewProperty("episodes")

	// The start date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	StartDate = schemaorg.NewProperty("startDate")

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an
	// event. Directors can be associated with individual items or with a series,
	// episode, clip.
	Director = schemaorg.NewProperty("director")

	// The production company or studio responsible for the item e.g. series, video
	// game, episode etc.
	ProductionCompany = schemaorg.NewProperty("productionCompany")

	// An episode of a tv, radio or game media within a series or season.
	Episode = schemaorg.NewProperty("episode")

	// Position of the season within an ordered group of seasons.
	SeasonNumber = schemaorg.NewProperty("seasonNumber")
)
