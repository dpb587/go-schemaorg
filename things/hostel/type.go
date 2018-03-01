package hostel

import "github.com/dpb587/go-schemaorg"

// // A hostel - cheap accommodation, often in shared dormitories.
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Hostel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
