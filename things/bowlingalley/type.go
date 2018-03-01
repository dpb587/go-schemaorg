package bowlingalley

import "github.com/dpb587/go-schemaorg"

// // A bowling alley.
var Type = schemaorg.NewDataType("http://schema.org", "BowlingAlley")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
