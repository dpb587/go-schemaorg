package broadcastchannel

import "github.com/dpb587/go-schemaorg"

// // A unique instance of a BroadcastService on a CableOrSatelliteService lineup.
var Type = schemaorg.NewDataType("http://schema.org", "BroadcastChannel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
