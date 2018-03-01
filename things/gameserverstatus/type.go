package gameserverstatus

import "github.com/dpb587/go-schemaorg"

// // Status of a game server.
var Type = schemaorg.NewDataType("http://schema.org", "GameServerStatus")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
