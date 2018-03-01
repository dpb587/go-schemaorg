package scheduleaction

import "github.com/dpb587/go-schemaorg"

// // <p>Scheduling future actions, events, or tasks.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/ReserveAction">ReserveAction</a>: Unlike
// ReserveAction, ScheduleAction allocates future actions (e.g. an event, a
// task, etc) towards a time slot / spatial allocation.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "ScheduleAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
