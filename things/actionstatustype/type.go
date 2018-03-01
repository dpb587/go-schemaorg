package actionstatustype

import "github.com/dpb587/go-schemaorg"

// // The status of an Action.
var Type = schemaorg.NewDataType("http://schema.org", "ActionStatusType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
