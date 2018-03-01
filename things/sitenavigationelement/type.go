package sitenavigationelement

import "github.com/dpb587/go-schemaorg"

// // A navigation element of the page.
var Type = schemaorg.NewDataType("http://schema.org", "SiteNavigationElement")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
