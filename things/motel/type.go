package motel

import "github.com/dpb587/go-schemaorg"

// // A motel.
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Motel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
