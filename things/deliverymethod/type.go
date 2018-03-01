package deliverymethod

import "github.com/dpb587/go-schemaorg"

// // <p>A delivery method is a standardized procedure for transferring the product
// or service to the destination of fulfillment chosen by the customer. Delivery
// methods are characterized by the means of transportation used, and by the
// organization or group that is the contracting party for the sending
// organization or person.</p>
// 
// <p>Commonly used values:</p>
// 
// <ul>
// <li>http://purl.org/goodrelations/v1#DeliveryModeDirectDownload</li>
// <li>http://purl.org/goodrelations/v1#DeliveryModeFreight</li>
// <li>http://purl.org/goodrelations/v1#DeliveryModeMail</li>
// <li>http://purl.org/goodrelations/v1#DeliveryModeOwnFleet</li>
// <li>http://purl.org/goodrelations/v1#DeliveryModePickUp</li>
// <li>http://purl.org/goodrelations/v1#DHL</li>
// <li>http://purl.org/goodrelations/v1#FederalExpress</li>
// <li>http://purl.org/goodrelations/v1#UPS</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "DeliveryMethod")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
