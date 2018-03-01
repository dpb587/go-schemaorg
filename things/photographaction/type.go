package photographaction

import "github.com/dpb587/go-schemaorg"

// // The act of capturing still images of objects using a camera.
var Type = schemaorg.NewDataType("http://schema.org", "PhotographAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
