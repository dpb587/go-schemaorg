package crematorium

import "github.com/dpb587/go-schemaorg"

// // A crematorium.
var Type = schemaorg.NewDataType("http://schema.org", "Crematorium")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
