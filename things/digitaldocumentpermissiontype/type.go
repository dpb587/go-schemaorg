package digitaldocumentpermissiontype

import "github.com/dpb587/go-schemaorg"

// // A type of permission which can be granted for accessing a digital document.
var Type = schemaorg.NewDataType("http://schema.org", "DigitalDocumentPermissionType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
