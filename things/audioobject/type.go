package audioobject

import "github.com/dpb587/go-schemaorg"

// // An audio file.
var Type = schemaorg.NewDataType("http://schema.org", "AudioObject")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
