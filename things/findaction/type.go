package findaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of finding an object.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/SearchAction">SearchAction</a>: FindAction is
// generally lead by a SearchAction, but not necessarily.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "FindAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
