package accommodation

import "github.com/dpb587/go-schemaorg"

// // An accommodation is a place that can accommodate human beings, e.g. a hotel
// room, a camping pitch, or a meeting room. Many accommodations are for
// overnight stays, but this is not a mandatory requirement.
// For more specific types of accommodations not defined in schema.org, one can
// use additionalType with external vocabularies.
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Accommodation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
