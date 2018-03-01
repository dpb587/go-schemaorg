package audience

import "github.com/dpb587/go-schemaorg"

// // Intended audience for an item, i.e. the group for whom the item was created.
var Type = schemaorg.NewDataType("http://schema.org", "Audience")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
