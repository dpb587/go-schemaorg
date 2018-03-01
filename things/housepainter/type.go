package housepainter

import "github.com/dpb587/go-schemaorg"

// // A house painting service.
var Type = schemaorg.NewDataType("http://schema.org", "HousePainter")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
