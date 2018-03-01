package achieveaction

import "github.com/dpb587/go-schemaorg"

// // The act of accomplishing something via previous efforts. It is an
// instantaneous action rather than an ongoing process.
var Type = schemaorg.NewDataType("http://schema.org", "AchieveAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
