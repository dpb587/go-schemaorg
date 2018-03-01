package videogameseries

import "github.com/dpb587/go-schemaorg"

// // A video game series.
var Type = schemaorg.NewDataType("http://schema.org", "VideoGameSeries")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
