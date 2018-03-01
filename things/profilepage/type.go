package profilepage

import "github.com/dpb587/go-schemaorg"

// // Web page type: Profile page.
var Type = schemaorg.NewDataType("http://schema.org", "ProfilePage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
