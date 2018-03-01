package resumeaction

import "github.com/dpb587/go-schemaorg"

// // The act of resuming a device or application which was formerly paused (e.g.
// resume music playback or resume a timer).
var Type = schemaorg.NewDataType("http://schema.org", "ResumeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
