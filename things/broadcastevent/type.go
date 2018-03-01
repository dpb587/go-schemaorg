package broadcastevent

import "github.com/dpb587/go-schemaorg"

// // An over the air or online broadcast event.
var Type = schemaorg.NewDataType("http://schema.org", "BroadcastEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
