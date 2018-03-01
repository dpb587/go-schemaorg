package replaceaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of object. The object that is being replaced.
	Replacee = schemaorg.NewProperty("replacee")

	// A sub property of object. The object that replaces.
	Replacer = schemaorg.NewProperty("replacer")
)
