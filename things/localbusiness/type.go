package localbusiness

import "github.com/dpb587/go-schemaorg"

// // A particular physical business or branch of an organization. Examples of
// LocalBusiness include a restaurant, a particular branch of a restaurant
// chain, a branch of a bank, a medical practice, a club, a bowling alley, etc.
var Type = schemaorg.NewDataType("http://schema.org", "LocalBusiness")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
