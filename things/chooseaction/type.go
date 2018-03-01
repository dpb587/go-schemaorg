package chooseaction

import "github.com/dpb587/go-schemaorg"

// // The act of expressing a preference from a set of options or a large or
// unbounded set of choices/options.
var Type = schemaorg.NewDataType("http://schema.org", "ChooseAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
