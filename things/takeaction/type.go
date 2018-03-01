package takeaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of gaining ownership of an object from an origin. Reciprocal of
// GiveAction.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/GiveAction">GiveAction</a>:
// The reciprocal of TakeAction.</li>
// <li><a class="localLink"
// href="http://schema.org/ReceiveAction">ReceiveAction</a>: Unlike
// ReceiveAction, TakeAction implies that ownership has been transfered.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "TakeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
