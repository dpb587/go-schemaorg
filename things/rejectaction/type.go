package rejectaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of rejecting to/adopting an object.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/AcceptAction">AcceptAction</a>: The antonym of
// RejectAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "RejectAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
