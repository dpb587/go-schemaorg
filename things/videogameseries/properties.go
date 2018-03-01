package videogameseries

import "github.com/dpb587/go-schemaorg"

var (
	// An item is an object within the game world that can be collected by a player
	// or, occasionally, a non-player character.
	GameItem = schemaorg.NewProperty("gameItem")

	// A season in a media series.
	Seasons = schemaorg.NewProperty("seasons")

	// The electronic systems used to play <a
	// href="http://en.wikipedia.org/wiki/Category:Video_game_platforms">video
	// games</a>.
	GamePlatform = schemaorg.NewProperty("gamePlatform")

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

	// Cheat codes to the game.
	CheatCode = schemaorg.NewProperty("cheatCode")

	// A piece of data that represents a particular aspect of a fictional character
	// (skill, power, character points, advantage, disadvantage).
	CharacterAttribute = schemaorg.NewProperty("characterAttribute")

	// Real or fictional location of the game (or part of game).
	GameLocation = schemaorg.NewProperty("gameLocation")

	// The composer of the soundtrack.
	MusicBy = schemaorg.NewProperty("musicBy")

	// A season that is part of the media series.
	ContainsSeason = schemaorg.NewProperty("containsSeason")

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can
	// be associated with individual items or with a series, episode, clip.
	Directors = schemaorg.NewProperty("directors")

	// Indicates whether this game is multi-player, co-op or single-player.  The
	// game can be marked as multi-player, co-op and single-player at the same time.
	PlayMode = schemaorg.NewProperty("playMode")

	// The number of seasons in this series.
	NumberOfSeasons = schemaorg.NewProperty("numberOfSeasons")

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an
	// event. Directors can be associated with individual items or with a series,
	// episode, clip.
	Director = schemaorg.NewProperty("director")

	// The production company or studio responsible for the item e.g. series, video
	// game, episode etc.
	ProductionCompany = schemaorg.NewProperty("productionCompany")

	// The task that a player-controlled character, or group of characters may
	// complete in order to gain a reward.
	Quest = schemaorg.NewProperty("quest")

	// Indicate how many people can play this game (minimum, maximum, or range).
	NumberOfPlayers = schemaorg.NewProperty("numberOfPlayers")

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated
	// with individual items or with a series, episode, clip.
	Actors = schemaorg.NewProperty("actors")

	// An episode of a tv, radio or game media within a series or season.
	Episode = schemaorg.NewProperty("episode")
)
