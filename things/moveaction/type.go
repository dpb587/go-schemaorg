package moveaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of an agent relocating to a place.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/TransferAction">TransferAction</a>: Unlike
// TransferAction, the subject of the move is a living Person or Organization
// rather than an inanimate object.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "MoveAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
