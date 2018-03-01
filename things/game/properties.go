package game

import "github.com/dpb587/go-schemaorg"

var (
	// An item is an object within the game world that can be collected by a player
	// or, occasionally, a non-player character.
	GameItem = schemaorg.NewProperty("gameItem")

	// A piece of data that represents a particular aspect of a fictional character
	// (skill, power, character points, advantage, disadvantage).
	CharacterAttribute = schemaorg.NewProperty("characterAttribute")

	// Real or fictional location of the game (or part of game).
	GameLocation = schemaorg.NewProperty("gameLocation")

	// The task that a player-controlled character, or group of characters may
	// complete in order to gain a reward.
	Quest = schemaorg.NewProperty("quest")

	// Indicate how many people can play this game (minimum, maximum, or range).
	NumberOfPlayers = schemaorg.NewProperty("numberOfPlayers")
)
