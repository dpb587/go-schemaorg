package ticket

import "github.com/dpb587/go-schemaorg"

// // Used to describe a ticket to an event, a flight, a bus ride, etc.
var Type = schemaorg.NewDataType("http://schema.org", "Ticket")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
