package brand

import "github.com/dpb587/go-schemaorg"

// // A brand is a name used by an organization or business person for labeling a
// product, product group, or similar.
var Type = schemaorg.NewDataType("http://schema.org", "Brand")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
