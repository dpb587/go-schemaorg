package howtotip

import "github.com/dpb587/go-schemaorg"

// // An explanation in the instructions for how to achieve a result. It provides
// supplementary information about a technique, supply, author's preference,
// etc. It can explain what could be done, or what should not be done, but
// doesn't specify what should be done (see HowToDirection).
var Type = schemaorg.NewDataType("http://schema.org", "HowToTip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
