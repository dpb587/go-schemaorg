package touristinformationcenter

import "github.com/dpb587/go-schemaorg"

// // A tourist information center.
var Type = schemaorg.NewDataType("http://schema.org", "TouristInformationCenter")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
