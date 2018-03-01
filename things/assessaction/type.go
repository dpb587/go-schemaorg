package assessaction

import "github.com/dpb587/go-schemaorg"

// // The act of forming one's opinion, reaction or sentiment.
var Type = schemaorg.NewDataType("http://schema.org", "AssessAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
