package reserveaction

import "github.com/dpb587/go-schemaorg"

// // <p>Reserving a concrete object.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/ScheduleAction">ScheduleAction</a></a>: Unlike
// ScheduleAction, ReserveAction reserves concrete objects (e.g. a table, a
// hotel) towards a time slot / spatial allocation.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "ReserveAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
