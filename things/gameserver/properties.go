package gameserver

import "github.com/dpb587/go-schemaorg"

var (
	// Video game which is played on this server.
	Game = schemaorg.NewProperty("game")

	// Number of players on the server.
	PlayersOnline = schemaorg.NewProperty("playersOnline")

	// Status of a game server.
	ServerStatus = schemaorg.NewProperty("serverStatus")
)
