package howtotool

import "github.com/dpb587/go-schemaorg"

// // A tool used (but not consumed) when performing instructions for how to
// achieve a result.
var Type = schemaorg.NewDataType("http://schema.org", "HowToTool")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
