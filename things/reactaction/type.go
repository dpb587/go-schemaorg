package reactaction

import "github.com/dpb587/go-schemaorg"

// // The act of responding instinctively and emotionally to an object, expressing
// a sentiment.
var Type = schemaorg.NewDataType("http://schema.org", "ReactAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
