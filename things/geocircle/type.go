package geocircle

import "github.com/dpb587/go-schemaorg"

// // A GeoCircle is a GeoShape representing a circular geographic area. As it is a
// GeoShape
//           it provides the simple textual property 'circle', but also allows
// the combination of postalCode alongside geoRadius.
//           The center of the circle can be indicated via the 'geoMidpoint'
// property, or more approximately using 'address', 'postalCode'.
var Type = schemaorg.NewDataType("http://schema.org", "GeoCircle")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
