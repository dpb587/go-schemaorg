package consumeaction

import "github.com/dpb587/go-schemaorg"

// // The act of ingesting information/resources/food.
var Type = schemaorg.NewDataType("http://schema.org", "ConsumeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
