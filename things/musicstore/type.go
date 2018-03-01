package musicstore

import "github.com/dpb587/go-schemaorg"

// // A music store.
var Type = schemaorg.NewDataType("http://schema.org", "MusicStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
