package mapcategorytype

import "github.com/dpb587/go-schemaorg"

// // An enumeration of several kinds of Map.
var Type = schemaorg.NewDataType("http://schema.org", "MapCategoryType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
