package befriendaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of forming a personal connection with someone (object)
// mutually/bidirectionally/symmetrically.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/FollowAction">FollowAction</a>: Unlike FollowAction,
// BefriendAction implies that the connection is reciprocal.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "BefriendAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
