package endorseaction

import "github.com/dpb587/go-schemaorg"

// // An agent approves/certifies/likes/supports/sanction an object.
var Type = schemaorg.NewDataType("http://schema.org", "EndorseAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
