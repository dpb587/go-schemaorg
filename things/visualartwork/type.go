package visualartwork

import "github.com/dpb587/go-schemaorg"

// // A work of art that is primarily visual in character.
var Type = schemaorg.NewDataType("http://schema.org", "VisualArtwork")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
