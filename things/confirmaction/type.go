package confirmaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of notifying someone that a future event/action is going to happen
// as expected.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/CancelAction">CancelAction</a>: The antonym of
// ConfirmAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "ConfirmAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
