package musicreleaseformattype

import "github.com/dpb587/go-schemaorg"

// // Format of this release (the type of recording media used, ie. compact disc,
// digital media, LP, etc.).
var Type = schemaorg.NewDataType("http://schema.org", "MusicReleaseFormatType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
