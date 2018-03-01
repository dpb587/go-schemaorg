package creditcard

import "github.com/dpb587/go-schemaorg"

// // <p>A card payment method of a particular brand or name.  Used to mark up a
// particular payment method and/or the financial product/service that supplies
// the card account.</p>
// 
// <p>Commonly used values:</p>
// 
// <ul>
// <li>http://purl.org/goodrelations/v1#AmericanExpress</li>
// <li>http://purl.org/goodrelations/v1#DinersClub</li>
// <li>http://purl.org/goodrelations/v1#Discover</li>
// <li>http://purl.org/goodrelations/v1#JCB</li>
// <li>http://purl.org/goodrelations/v1#MasterCard</li>
// <li>http://purl.org/goodrelations/v1#VISA</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "CreditCard")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
