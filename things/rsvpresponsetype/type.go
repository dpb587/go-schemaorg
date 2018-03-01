package rsvpresponsetype

import "github.com/dpb587/go-schemaorg"

// // RsvpResponseType is an enumeration type whose instances represent responding
// to an RSVP request.
var Type = schemaorg.NewDataType("http://schema.org", "RsvpResponseType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
