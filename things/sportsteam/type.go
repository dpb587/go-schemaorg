package sportsteam

import "github.com/dpb587/go-schemaorg"

// // Organization: Sports team.
var Type = schemaorg.NewDataType("http://schema.org", "SportsTeam")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
