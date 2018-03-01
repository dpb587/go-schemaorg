package geocircle

import "github.com/dpb587/go-schemaorg"

var (
	// Indicates the approximate radius of a GeoCircle (metres unless indicated
	// otherwise via Distance notation).
	GeoRadius = schemaorg.NewProperty("geoRadius")

	// Indicates the GeoCoordinates at the centre of a GeoShape e.g. GeoCircle.
	GeoMidpoint = schemaorg.NewProperty("geoMidpoint")
)
