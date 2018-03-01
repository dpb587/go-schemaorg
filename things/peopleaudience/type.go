package peopleaudience

import "github.com/dpb587/go-schemaorg"

// // A set of characteristics belonging to people, e.g. who compose an item's
// target audience.
var Type = schemaorg.NewDataType("http://schema.org", "PeopleAudience")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
