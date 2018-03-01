package menuitem

import "github.com/dpb587/go-schemaorg"

// // A food or drink item listed in a menu or menu section.
var Type = schemaorg.NewDataType("http://schema.org", "MenuItem")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
