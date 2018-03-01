package pawnshop

import "github.com/dpb587/go-schemaorg"

// // A shop that will buy, or lend money against the security of, personal
// possessions.
var Type = schemaorg.NewDataType("http://schema.org", "PawnShop")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
