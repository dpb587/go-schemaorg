package performaction

import "github.com/dpb587/go-schemaorg"

// // The act of participating in performance arts.
var Type = schemaorg.NewDataType("http://schema.org", "PerformAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
