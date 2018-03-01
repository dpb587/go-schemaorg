package rating

import "github.com/dpb587/go-schemaorg"

// // A rating is an evaluation on a numeric scale, such as 1 to 5 stars.
var Type = schemaorg.NewDataType("http://schema.org", "Rating")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
