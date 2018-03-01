package cancelaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of asserting that a future event/action is no longer going to
// happen.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/ConfirmAction">ConfirmAction</a>: The antonym of
// CancelAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "CancelAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
