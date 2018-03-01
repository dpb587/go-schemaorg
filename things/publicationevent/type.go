package publicationevent

import "github.com/dpb587/go-schemaorg"

// // A PublicationEvent corresponds indifferently to the event of publication for
// a CreativeWork of any type e.g. a broadcast event, an on-demand event, a
// book/journal publication via a variety of delivery media.
var Type = schemaorg.NewDataType("http://schema.org", "PublicationEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
