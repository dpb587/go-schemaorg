package checkoutaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of an agent communicating (service provider, social media, etc)
// their departure of a previously reserved service (e.g. flight check in) or
// place (e.g. hotel).</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/CheckInAction">CheckInAction</a>: The antonym of
// CheckOutAction.</li>
// <li><a class="localLink"
// href="http://schema.org/DepartAction">DepartAction</a>: Unlike DepartAction,
// CheckOutAction implies that the agent is informing/confirming the end of a
// previously reserved service.</li>
// <li><a class="localLink"
// href="http://schema.org/CancelAction">CancelAction</a>: Unlike CancelAction,
// CheckOutAction implies that the agent is informing/confirming the end of a
// previously reserved service.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "CheckOutAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
