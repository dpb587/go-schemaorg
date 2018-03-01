package receiveaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of physically/electronically taking delivery of an object thathas
// been transferred from an origin to a destination. Reciprocal of
// SendAction.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/SendAction">SendAction</a>:
// The reciprocal of ReceiveAction.</li>
// <li><a class="localLink" href="http://schema.org/TakeAction">TakeAction</a>:
// Unlike TakeAction, ReceiveAction does not imply that the ownership has been
// transfered (e.g. I can receive a package, but it does not mean the package is
// now mine).</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "ReceiveAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
