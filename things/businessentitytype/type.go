package businessentitytype

import "github.com/dpb587/go-schemaorg"

// // <p>A business entity type is a conceptual entity representing the legal form,
// the size, the main line of business, the position in the value chain, or any
// combination thereof, of an organization or business person.</p>
// 
// <p>Commonly used values:</p>
// 
// <ul>
// <li>http://purl.org/goodrelations/v1#Business</li>
// <li>http://purl.org/goodrelations/v1#Enduser</li>
// <li>http://purl.org/goodrelations/v1#PublicInstitution</li>
// <li>http://purl.org/goodrelations/v1#Reseller</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "BusinessEntityType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
