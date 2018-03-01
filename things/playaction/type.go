package playaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of playing/exercising/training/performing for enjoyment, leisure,
// recreation, Competition or exercise.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/ListenAction">ListenAction</a>: Unlike ListenAction
// (which is under ConsumeAction), PlayAction refers to performing for an
// audience or at an event, rather than consuming music.</li>
// <li><a class="localLink"
// href="http://schema.org/WatchAction">WatchAction</a>: Unlike WatchAction
// (which is under ConsumeAction), PlayAction refers to showing/displaying for
// an audience or at an event, rather than consuming visual content.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "PlayAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
