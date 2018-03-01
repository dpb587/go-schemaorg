package videogame

import "github.com/dpb587/go-schemaorg"

var (
	// The electronic systems used to play <a
	// href="http://en.wikipedia.org/wiki/Category:Video_game_platforms">video
	// games</a>.
	GamePlatform = schemaorg.NewProperty("gamePlatform")

	// The server on which  it is possible to play the game.
	GameServer = schemaorg.NewProperty("gameServer")

	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors
	// can be associated with individual items or with a series, episode, clip.
	Actor = schemaorg.NewProperty("actor")

	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer = schemaorg.NewProperty("trailer")

	// Cheat codes to the game.
	CheatCode = schemaorg.NewProperty("cheatCode")

	// Links to tips, tactics, etc.
	GameTip = schemaorg.NewProperty("gameTip")

	// The composer of the soundtrack.
	MusicBy = schemaorg.NewProperty("musicBy")

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can
	// be associated with individual items or with a series, episode, clip.
	Directors = schemaorg.NewProperty("directors")

	// Indicates whether this game is multi-player, co-op or single-player.  The
	// game can be marked as multi-player, co-op and single-player at the same time.
	PlayMode = schemaorg.NewProperty("playMode")

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an
	// event. Directors can be associated with individual items or with a series,
	// episode, clip.
	Director = schemaorg.NewProperty("director")

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated
	// with individual items or with a series, episode, clip.
	Actors = schemaorg.NewProperty("actors")
)
