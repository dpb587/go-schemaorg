package joinaction

import "github.com/dpb587/go-schemaorg"

// // <p>An agent joins an event/group with participants/friends at a location.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/RegisterAction">RegisterAction</a>: Unlike
// RegisterAction, JoinAction refers to joining a group/team of people.</li>
// <li><a class="localLink"
// href="http://schema.org/SubscribeAction">SubscribeAction</a>: Unlike
// SubscribeAction, JoinAction does not imply that you'll be receiving
// updates.</li>
// <li><a class="localLink"
// href="http://schema.org/FollowAction">FollowAction</a>: Unlike FollowAction,
// JoinAction does not imply that you'll be polling for updates.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "JoinAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
