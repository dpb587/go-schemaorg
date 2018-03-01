package tieaction

import "github.com/dpb587/go-schemaorg"

// // The act of reaching a draw in a competitive activity.
var Type = schemaorg.NewDataType("http://schema.org", "TieAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
