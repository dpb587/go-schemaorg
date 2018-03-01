package askaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of posing a question / favor to someone.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/ReplyAction">ReplyAction</a>: Appears generally as a
// response to AskAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "AskAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
