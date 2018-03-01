package warrantyscope

import "github.com/dpb587/go-schemaorg"

// // <p>A range of of services that will be provided to a customer free of charge
// in case of a defect or malfunction of a product.</p>
// 
// <p>Commonly used values:</p>
// 
// <ul>
// <li>http://purl.org/goodrelations/v1#Labor-BringIn</li>
// <li>http://purl.org/goodrelations/v1#PartsAndLabor-BringIn</li>
// <li>http://purl.org/goodrelations/v1#PartsAndLabor-PickUp</li>
// </ul>
// 
var Type = schemaorg.NewDataType("http://schema.org", "WarrantyScope")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
