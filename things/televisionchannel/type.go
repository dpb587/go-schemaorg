package televisionchannel

import "github.com/dpb587/go-schemaorg"

// // A unique instance of a television BroadcastService on a
// CableOrSatelliteService lineup.
var Type = schemaorg.NewDataType("http://schema.org", "TelevisionChannel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
