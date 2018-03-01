package listitem

import "github.com/dpb587/go-schemaorg"

// // An list item, e.g. a step in a checklist or how-to description.
var Type = schemaorg.NewDataType("http://schema.org", "ListItem")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
