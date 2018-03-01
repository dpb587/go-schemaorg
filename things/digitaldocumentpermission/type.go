package digitaldocumentpermission

import "github.com/dpb587/go-schemaorg"

// // A permission for a particular person or group to access a particular file.
var Type = schemaorg.NewDataType("http://schema.org", "DigitalDocumentPermission")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
