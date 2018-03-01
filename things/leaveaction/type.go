package leaveaction

import "github.com/dpb587/go-schemaorg"

// // <p>An agent leaves an event / group with participants/friends at a
// location.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/JoinAction">JoinAction</a>:
// The antonym of LeaveAction.</li>
// <li><a class="localLink"
// href="http://schema.org/UnRegisterAction">UnRegisterAction</a>: Unlike
// UnRegisterAction, LeaveAction implies leaving a group/team of people rather
// than a service.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "LeaveAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
