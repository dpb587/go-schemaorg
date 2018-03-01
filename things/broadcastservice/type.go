package broadcastservice

import "github.com/dpb587/go-schemaorg"

// // A delivery service through which content is provided via broadcast over the
// air or online.
var Type = schemaorg.NewDataType("http://schema.org", "BroadcastService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
