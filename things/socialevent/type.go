package socialevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Social event.
var Type = schemaorg.NewDataType("http://schema.org", "SocialEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
