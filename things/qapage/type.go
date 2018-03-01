package qapage

import "github.com/dpb587/go-schemaorg"

// // A QAPage is a WebPage focussed on a specific Question and its Answer(s), e.g.
// in a question answering site or documenting Frequently Asked Questions
// (FAQs).
var Type = schemaorg.NewDataType("http://schema.org", "QAPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
