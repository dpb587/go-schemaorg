package contactpointoption

import "github.com/dpb587/go-schemaorg"

// // Enumerated options related to a ContactPoint.
var Type = schemaorg.NewDataType("http://schema.org", "ContactPointOption")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
