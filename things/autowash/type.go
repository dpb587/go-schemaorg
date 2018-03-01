package autowash

import "github.com/dpb587/go-schemaorg"

// // A car wash business.
var Type = schemaorg.NewDataType("http://schema.org", "AutoWash")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
