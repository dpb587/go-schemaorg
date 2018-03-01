package apartment

import "github.com/dpb587/go-schemaorg"

// // An apartment (in American English) or flat (in British English) is a
// self-contained housing unit (a type of residential real estate) that occupies
// only part of a building (Source: Wikipedia, the free encyclopedia, see <a
// href="http://en.wikipedia.org/wiki/Apartment">http://en.wikipedia.org/wiki/Apartment</a>).
var Type = schemaorg.NewDataType("http://schema.org", "Apartment")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
