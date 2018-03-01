package lakebodyofwater

import "github.com/dpb587/go-schemaorg"

// // A lake (for example, Lake Pontrachain).
var Type = schemaorg.NewDataType("http://schema.org", "LakeBodyOfWater")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
