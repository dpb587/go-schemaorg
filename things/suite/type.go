package suite

import "github.com/dpb587/go-schemaorg"

// // A suite in a hotel or other public accommodation, denotes a class of luxury
// accommodations, the key feature of which is multiple rooms (Source:
// Wikipedia, the free encyclopedia, see <a
// href="http://en.wikipedia.org/wiki/Suite_(hotel)">http://en.wikipedia.org/wiki/Suite_(hotel)</a>).
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Suite")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
