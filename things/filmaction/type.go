package filmaction

import "github.com/dpb587/go-schemaorg"

// // The act of capturing sound and moving images on film, video, or digitally.
var Type = schemaorg.NewDataType("http://schema.org", "FilmAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
