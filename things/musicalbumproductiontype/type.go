package musicalbumproductiontype

import "github.com/dpb587/go-schemaorg"

// // Classification of the album by it's type of content: soundtrack, live album,
// studio album, etc.
var Type = schemaorg.NewDataType("http://schema.org", "MusicAlbumProductionType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
