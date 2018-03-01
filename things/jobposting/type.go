package jobposting

import "github.com/dpb587/go-schemaorg"

// // A listing that describes a job opening in a certain organization.
var Type = schemaorg.NewDataType("http://schema.org", "JobPosting")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
