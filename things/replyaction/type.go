package replyaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of responding to a question/message asked/sent by the object.
// Related to <a class="localLink"
// href="http://schema.org/AskAction">AskAction</a></p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink" href="http://schema.org/AskAction">AskAction</a>:
// Appears generally as an origin of a ReplyAction.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "ReplyAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
