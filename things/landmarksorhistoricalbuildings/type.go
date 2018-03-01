package landmarksorhistoricalbuildings

import "github.com/dpb587/go-schemaorg"

// // An historical landmark or building.
var Type = schemaorg.NewDataType("http://schema.org", "LandmarksOrHistoricalBuildings")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
