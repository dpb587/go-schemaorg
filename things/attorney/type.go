package attorney

import "github.com/dpb587/go-schemaorg"

// // Professional service: Attorney. </p>
// 
// <p>This type is deprecated - <a class="localLink"
// href="http://schema.org/LegalService">LegalService</a> is more inclusive and
// less ambiguous.
var Type = schemaorg.NewDataType("http://schema.org", "Attorney")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
