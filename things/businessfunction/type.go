package businessfunction

import "github.com/dpb587/go-schemaorg"

// // <p>The business function specifies the type of activity or access (i.e., the
// bundle of rights) offered by the organization or business person through the
// offer. Typical are sell, rental or lease, maintenance or repair, manufacture
// / produce, recycle / dispose, engineering / construction, or installation.
// Proprietary specifications of access rights are also instances of this
// class.</p>
// 
// <p>Commonly used values:</p>
// 
// <ul>
// <li>http://purl.org/goodrelations/v1#ConstructionInstallation</li>
// <li>http://purl.org/goodrelations/v1#Dispose</li>
// <li>http://purl.org/goodrelations/v1#LeaseOut</li>
// <li>http://purl.org/goodrelations/v1#Maintain</li>
// <li>http://purl.org/goodrelations/v1#ProvideService</li>
// <li>http://purl.org/goodrelations/v1#Repair</li>
// <li>http://purl.org/goodrelations/v1#Sell</li>
// <li>http://purl.org/goodrelations/v1#Buy</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "BusinessFunction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
