package mobilephonestore

import "github.com/dpb587/go-schemaorg"

// // A store that sells mobile phones and related accessories.
var Type = schemaorg.NewDataType("http://schema.org", "MobilePhoneStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
