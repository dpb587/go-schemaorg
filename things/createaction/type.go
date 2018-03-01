package createaction

import "github.com/dpb587/go-schemaorg"

// // The act of deliberately creating/producing/generating/building a result out
// of the agent.
var Type = schemaorg.NewDataType("http://schema.org", "CreateAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
