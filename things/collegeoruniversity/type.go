package collegeoruniversity

import "github.com/dpb587/go-schemaorg"

// // A college, university, or other third-level educational institution.
var Type = schemaorg.NewDataType("http://schema.org", "CollegeOrUniversity")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
