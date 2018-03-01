package planaction

import "github.com/dpb587/go-schemaorg"

// // The act of planning the execution of an event/task/action/reservation/plan to
// a future date.
var Type = schemaorg.NewDataType("http://schema.org", "PlanAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
