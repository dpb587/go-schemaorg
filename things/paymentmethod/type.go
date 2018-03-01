package paymentmethod

import "github.com/dpb587/go-schemaorg"

// // <p>A payment method is a standardized procedure for transferring the monetary
// amount for a purchase. Payment methods are characterized by the legal and
// technical structures used, and by the organization or group carrying out the
// transaction.</p>
// 
// <p>Commonly used values:</p>
// 
// <ul>
// <li>http://purl.org/goodrelations/v1#ByBankTransferInAdvance</li>
// <li>http://purl.org/goodrelations/v1#ByInvoice</li>
// <li>http://purl.org/goodrelations/v1#Cash</li>
// <li>http://purl.org/goodrelations/v1#CheckInAdvance</li>
// <li>http://purl.org/goodrelations/v1#COD</li>
// <li>http://purl.org/goodrelations/v1#DirectDebit</li>
// <li>http://purl.org/goodrelations/v1#GoogleCheckout</li>
// <li>http://purl.org/goodrelations/v1#PayPal</li>
// <li>http://purl.org/goodrelations/v1#PaySwarm</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "PaymentMethod")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
