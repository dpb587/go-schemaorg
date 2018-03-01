package inviteaction

import "github.com/dpb587/go-schemaorg"

// // The act of asking someone to attend an event. Reciprocal of RsvpAction.
var Type = schemaorg.NewDataType("http://schema.org", "InviteAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
