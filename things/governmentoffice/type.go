package governmentoffice

import "github.com/dpb587/go-schemaorg"

// // A government office&#x2014;for example, an IRS or DMV office.
var Type = schemaorg.NewDataType("http://schema.org", "GovernmentOffice")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
