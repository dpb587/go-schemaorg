package event

import "github.com/dpb587/go-schemaorg"

// // An event happening at a certain time and location, such as a concert,
// lecture, or festival. Ticketing information may be added via the <a
// class="localLink" href="http://schema.org/offers">offers</a> property.
// Repeated events may be structured as separate Event objects.
var Type = schemaorg.NewDataType("http://schema.org", "Event")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
