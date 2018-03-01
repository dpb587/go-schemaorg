package rsvpaction

import "github.com/dpb587/go-schemaorg"

// // The act of notifying an event organizer as to whether you expect to attend
// the event.
var Type = schemaorg.NewDataType("http://schema.org", "RsvpAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
