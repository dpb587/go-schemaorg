package registeraction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of registering to be a user of a service, product or web page.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/JoinAction">JoinAction</a>:
// Unlike JoinAction, RegisterAction implies you are registering to be a user of
// a service, <em>not</em> a group/team of people.</li>
// <li>[FollowAction]]: Unlike FollowAction, RegisterAction doesn't imply that
// the agent is expecting to poll for updates from the object.</li>
// <li><a class="localLink"
// href="http://schema.org/SubscribeAction">SubscribeAction</a>: Unlike
// SubscribeAction, RegisterAction doesn't imply that the agent is expecting
// updates from the object.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "RegisterAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
