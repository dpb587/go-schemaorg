package downloadaction

import "github.com/dpb587/go-schemaorg"

// // The act of downloading an object.
var Type = schemaorg.NewDataType("http://schema.org", "DownloadAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
