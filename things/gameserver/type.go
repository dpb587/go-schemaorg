package gameserver

import "github.com/dpb587/go-schemaorg"

// // Server that provides game interaction in a multiplayer game.
var Type = schemaorg.NewDataType("http://schema.org", "GameServer")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
