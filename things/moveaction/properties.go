package moveaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of location. The final location of the object or the agent
	// after the action.
	ToLocation = schemaorg.NewProperty("toLocation")

	// A sub property of location. The original location of the object or the agent
	// before the action.
	FromLocation = schemaorg.NewProperty("fromLocation")
)
