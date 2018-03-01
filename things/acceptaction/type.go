package acceptaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of committing to/adopting an object.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/RejectAction">RejectAction</a>: The antonym of
// AcceptAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "AcceptAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
