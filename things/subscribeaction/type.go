package subscribeaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of forming a personal connection with someone/something (object)
// unidirectionally/asymmetrically to get updates pushed to.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/FollowAction">FollowAction</a>: Unlike FollowAction,
// SubscribeAction implies that the subscriber acts as a passive agent being
// constantly/actively pushed for updates.</li>
// <li><a class="localLink"
// href="http://schema.org/RegisterAction">RegisterAction</a>: Unlike
// RegisterAction, SubscribeAction implies that the agent is interested in
// continuing receiving updates from the object.</li>
// <li><a class="localLink" href="http://schema.org/JoinAction">JoinAction</a>:
// Unlike JoinAction, SubscribeAction implies that the agent is interested in
// continuing receiving updates from the object.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "SubscribeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
