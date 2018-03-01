package gameplaymode

import "github.com/dpb587/go-schemaorg"

// // Indicates whether this game is multi-player, co-op or single-player.
var Type = schemaorg.NewDataType("http://schema.org", "GamePlayMode")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
