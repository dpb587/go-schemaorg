package departmentstore

import "github.com/dpb587/go-schemaorg"

// // A department store.
var Type = schemaorg.NewDataType("http://schema.org", "DepartmentStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
