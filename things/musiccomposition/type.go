package musiccomposition

import "github.com/dpb587/go-schemaorg"

// // A musical composition.
var Type = schemaorg.NewDataType("http://schema.org", "MusicComposition")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
