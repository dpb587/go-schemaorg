package lendaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of providing an object under an agreement that it will be returned
// at a later date. Reciprocal of BorrowAction.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/BorrowAction">BorrowAction</a>: Reciprocal of
// LendAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "LendAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
