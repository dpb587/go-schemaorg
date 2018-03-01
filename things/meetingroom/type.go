package meetingroom

import "github.com/dpb587/go-schemaorg"

// // A meeting room, conference room, or conference hall is a room provided for
// singular events such as business conferences and meetings (Source: Wikipedia,
// the free encyclopedia, see <a
// href="http://en.wikipedia.org/wiki/Conference_hall">http://en.wikipedia.org/wiki/Conference_hall</a>).
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "MeetingRoom")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
