package musicgroup

import "github.com/dpb587/go-schemaorg"

// // A musical group, such as a band, an orchestra, or a choir. Can also be a solo
// musician.
var Type = schemaorg.NewDataType("http://schema.org", "MusicGroup")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
