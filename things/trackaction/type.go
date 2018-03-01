package trackaction

import "github.com/dpb587/go-schemaorg"

// // <p>An agent tracks an object for updates.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/FollowAction">FollowAction</a>: Unlike FollowAction,
// TrackAction refers to the interest on the location of innanimates
// objects.</li>
// <li><a class="localLink"
// href="http://schema.org/SubscribeAction">SubscribeAction</a>: Unlike
// SubscribeAction, TrackAction refers to  the interest on the location of
// innanimate objects.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "TrackAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
