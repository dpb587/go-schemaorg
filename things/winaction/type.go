package winaction

import "github.com/dpb587/go-schemaorg"

// // The act of achieving victory in a competitive activity.
var Type = schemaorg.NewDataType("http://schema.org", "WinAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
