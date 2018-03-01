package beautysalon

import "github.com/dpb587/go-schemaorg"

// // Beauty salon.
var Type = schemaorg.NewDataType("http://schema.org", "BeautySalon")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
