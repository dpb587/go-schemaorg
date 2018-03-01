package unregisteraction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of un-registering from a service.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/RegisterAction">RegisterAction</a>: antonym of
// UnRegisterAction.</li>
// <li><a class="localLink" href="http://schema.org/Leave">Leave</a>: Unlike
// LeaveAction, UnRegisterAction implies that you are unregistering from a
// service you werer previously registered, rather than leaving a team/group of
// people.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "UnRegisterAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
