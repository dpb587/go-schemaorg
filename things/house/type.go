package house

import "github.com/dpb587/go-schemaorg"

// // A house is a building or structure that has the ability to be occupied for
// habitation by humans or other creatures (Source: Wikipedia, the free
// encyclopedia, see <a
// href="http://en.wikipedia.org/wiki/House">http://en.wikipedia.org/wiki/House</a>).
var Type = schemaorg.NewDataType("http://schema.org", "House")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
