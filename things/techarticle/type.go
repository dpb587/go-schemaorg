package techarticle

import "github.com/dpb587/go-schemaorg"

// // A technical article - Example: How-to (task) topics, step-by-step, procedural
// troubleshooting, specifications, etc.
var Type = schemaorg.NewDataType("http://schema.org", "TechArticle")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
