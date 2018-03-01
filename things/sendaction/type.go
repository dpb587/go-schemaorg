package sendaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of physically/electronically dispatching an object for transfer
// from an origin to a destination.Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/ReceiveAction">ReceiveAction</a>: The reciprocal of
// SendAction.</li>
// <li><a class="localLink" href="http://schema.org/GiveAction">GiveAction</a>:
// Unlike GiveAction, SendAction does not imply the transfer of ownership (e.g.
// I can send you my laptop, but I'm not necessarily giving it to you).</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "SendAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
