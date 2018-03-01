package webpage

import "github.com/dpb587/go-schemaorg"

// // A web page. Every web page is implicitly assumed to be declared to be of type
// WebPage, so the various properties about that webpage, such as
// <code>breadcrumb</code> may be used. We recommend explicit declaration if
// these properties are specified, but if they are found outside of an
// itemscope, they will be assumed to be about the page.
var Type = schemaorg.NewDataType("http://schema.org", "WebPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
