package room

import "github.com/dpb587/go-schemaorg"

// // A room is a distinguishable space within a structure, usually separated from
// other spaces by interior walls. (Source: Wikipedia, the free encyclopedia,
// see <a
// href="http://en.wikipedia.org/wiki/Room">http://en.wikipedia.org/wiki/Room</a>).
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Room")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
