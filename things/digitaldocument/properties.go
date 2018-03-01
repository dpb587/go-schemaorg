package digitaldocument

import "github.com/dpb587/go-schemaorg"

var (
	// A permission related to the access to this document (e.g. permission to read
	// or write an electronic document). For a public document, specify a grantee
	// with an Audience with audienceType equal to "public".
	HasDigitalDocumentPermission = schemaorg.NewProperty("hasDigitalDocumentPermission")
)
