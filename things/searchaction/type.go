package searchaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of searching for an object.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/FindAction">FindAction</a>:
// SearchAction generally leads to a FindAction, but not necessarily.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "SearchAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
