package parcelservice

import "github.com/dpb587/go-schemaorg"

// // <p>A private parcel service as the delivery mode available for a certain
// offer.</p>
// 
// <p>Commonly used values:</p>
// 
// <ul>
// <li>http://purl.org/goodrelations/v1#DHL</li>
// <li>http://purl.org/goodrelations/v1#FederalExpress</li>
// <li>http://purl.org/goodrelations/v1#UPS</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "ParcelService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
