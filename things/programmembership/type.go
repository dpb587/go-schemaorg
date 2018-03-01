package programmembership

import "github.com/dpb587/go-schemaorg"

// // Used to describe membership in a loyalty programs (e.g. "StarAliance"),
// traveler clubs (e.g. "AAA"), purchase clubs ("Safeway Club"), etc.
var Type = schemaorg.NewDataType("http://schema.org", "ProgramMembership")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
