package radiochannel

import "github.com/dpb587/go-schemaorg"

// // A unique instance of a radio BroadcastService on a CableOrSatelliteService
// lineup.
var Type = schemaorg.NewDataType("http://schema.org", "RadioChannel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
