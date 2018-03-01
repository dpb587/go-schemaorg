package watchaction

import "github.com/dpb587/go-schemaorg"

// // The act of consuming dynamic/moving visual content.
var Type = schemaorg.NewDataType("http://schema.org", "WatchAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
