package healthandbeautybusiness

import "github.com/dpb587/go-schemaorg"

// // Health and beauty.
var Type = schemaorg.NewDataType("http://schema.org", "HealthAndBeautyBusiness")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
