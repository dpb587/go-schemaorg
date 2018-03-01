package giveaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of transferring ownership of an object to a destination.
// Reciprocal of TakeAction.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/TakeAction">TakeAction</a>:
// Reciprocal of GiveAction.</li>
// <li><a class="localLink" href="http://schema.org/SendAction">SendAction</a>:
// Unlike SendAction, GiveAction implies that ownership is being transferred
// (e.g. I may send my laptop to you, but that doesn't mean I'm giving it to
// you).</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "GiveAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
