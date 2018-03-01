package suspendaction

import "github.com/dpb587/go-schemaorg"

// // The act of momentarily pausing a device or application (e.g. pause music
// playback or pause a timer).
var Type = schemaorg.NewDataType("http://schema.org", "SuspendAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
