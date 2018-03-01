package applyaction

import "github.com/dpb587/go-schemaorg"

// // <p>The act of registering to an organization/service without the guarantee to
// receive it.</p>
// 
// <p>Related actions:</p>
// 
// <ul>
// <li><a class="localLink"
// href="http://schema.org/RegisterAction">RegisterAction</a>: Unlike
// RegisterAction, ApplyAction has no guarantees that the application will be
// accepted.</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "ApplyAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
