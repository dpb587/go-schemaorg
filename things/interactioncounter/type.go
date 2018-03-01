package interactioncounter

import "github.com/dpb587/go-schemaorg"

// // A summary of how users have interacted with this CreativeWork. In most cases,
// authors will use a subtype to specify the specific type of interaction.
var Type = schemaorg.NewDataType("http://schema.org", "InteractionCounter")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
