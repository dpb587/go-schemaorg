package digitaldocumentpermission

import "github.com/dpb587/go-schemaorg"

var (
	// The person, organization, contact point, or audience that has been granted
	// this permission.
	Grantee = schemaorg.NewProperty("grantee")

	// The type of permission granted the person, organization, or audience.
	PermissionType = schemaorg.NewProperty("permissionType")
)
