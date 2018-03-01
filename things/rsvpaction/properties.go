package rsvpaction

import "github.com/dpb587/go-schemaorg"

var (
	// If responding yes, the number of guests who will attend in addition to the
	// invitee.
	AdditionalNumberOfGuests = schemaorg.NewProperty("additionalNumberOfGuests")

	// Comments, typically from users.
	Comment = schemaorg.NewProperty("comment")

	// The response (yes, no, maybe) to the RSVP.
	RsvpResponse = schemaorg.NewProperty("rsvpResponse")
)
