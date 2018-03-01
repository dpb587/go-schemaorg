package openinghoursspecification

import "github.com/dpb587/go-schemaorg"

// // A structured value providing information about the opening hours of a place
// or a certain service inside a place.</p>
// 
// <p>The place is <strong>open</strong> if the <a class="localLink"
// href="http://schema.org/opens">opens</a> property is specified, and
// <strong>closed</strong> otherwise.</p>
// 
// <p>If the value for the <a class="localLink"
// href="http://schema.org/closes">closes</a> property is less than the value
// for the <a class="localLink" href="http://schema.org/opens">opens</a>
// property then the hour range is assumed to span over the next day.
var Type = schemaorg.NewDataType("http://schema.org", "OpeningHoursSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
