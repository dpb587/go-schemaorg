package landform

import "github.com/dpb587/go-schemaorg"

// // A landform or physical feature.  Landform elements include mountains, plains,
// lakes, rivers, seascape and oceanic waterbody interface features such as
// bays, peninsulas, seas and so forth, including sub-aqueous terrain features
// such as submersed mountain ranges, volcanoes, and the great ocean basins.
var Type = schemaorg.NewDataType("http://schema.org", "Landform")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
