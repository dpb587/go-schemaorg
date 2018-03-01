package borrowaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of obtaining an object under an agreement to return it at a later
// date. Reciprocal of LendAction.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/LendAction">LendAction</a>:
// Reciprocal of BorrowAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "BorrowAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
