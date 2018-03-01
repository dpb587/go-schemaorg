package campingpitch

import "github.com/dpb587/go-schemaorg"

// // A camping pitch is an individual place for overnight stay in the outdoors,
// typically being part of a larger camping site.
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "CampingPitch")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
